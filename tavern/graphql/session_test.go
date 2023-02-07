package graphql_test

import (
	"context"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kcarretto/realm/tavern/ent"
	"github.com/kcarretto/realm/tavern/ent/enttest"
	"github.com/kcarretto/realm/tavern/ent/session"
	"github.com/kcarretto/realm/tavern/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClaimTasks(t *testing.T) {
	// Initialize Test Context
	ctx := context.Background()

	// Initialize DB Backend
	graph := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer graph.Close()

	// Initialize sample data
	testSessions := []*ent.Session{
		graph.Session.Create().
			SetPrincipal("admin").
			SetAgentIdentifier("TEST").
			SetIdentifier("SOME_ID").
			SetHostIdentifier("SOME_HOST_ID").
			SetHostname("SOME_HOSTNAME").
			SetLastSeenAt(time.Now().Add(-10 * time.Minute)).
			SaveX(ctx),
	}
	testTome := graph.Tome.Create().
		SetName("Test Tome").
		SetDescription("Ensures the world feels greeted").
		SetEldritch(`print("Hello World!")`).
		SaveX(ctx)
	testJob := graph.Job.Create().
		SetName("howdy-ho").
		SetTome(testTome).
		SaveX(ctx)
	testTasks := []*ent.Task{
		graph.Task.Create().
			SetJob(testJob).
			SetSession(testSessions[0]).
			SaveX(ctx),
	}

	// Create a new GraphQL server (needed for auth middleware)
	srv := handler.NewDefaultServer(graphql.NewSchema(graph))

	// Create a new GraphQL client (connected to our http server)
	gqlClient := client.New(srv)

	// Define the ClaimTasks mutation
	mut := `
mutation newClaimTasksTest($input: ClaimTasksInput!) {
	claimTasks(input: $input) {
		id 
	}
}
`

	// Create a closure to execute the mutation
	claimTasks := func(input map[string]any) ([]int, error) {
		// Make our request to the GraphQL API
		var resp struct {
			ClaimTasks []struct{ ID string }
		}
		err := gqlClient.Post(mut, &resp, client.Var("input", input))
		if err != nil {
			return nil, err
		}

		// Parse Task IDs from response
		ids := make([]int, 0, len(resp.ClaimTasks))
		for _, task := range resp.ClaimTasks {
			ids = append(ids, convertID(task.ID))
		}
		return ids, nil
	}

	/*
	 * Test when the `claimTasks` mutation is run by a session that does not exist.
	 *
	 * Expected that the session is created, but no tasks are returned.
	 */
	t.Run("NewSession", func(t *testing.T) {
		expectedIdentifier := "NEW_ID"
		expected := map[string]any{
			"principal":         "newuser",
			"hostname":          "NEW_HOSTNAME",
			"sessionIdentifier": expectedIdentifier,
			"hostIdentifier":    "NEW_HOST_ID",
			"agentIdentifier":   "NEW_AGENT_ID",
		}
		ids, err := claimTasks(expected)
		require.NoError(t, err)
		assert.Empty(t, ids)
		testSession, err := graph.Session.Query().
			Where(session.Identifier(expectedIdentifier)).
			Only(ctx)
		require.NoError(t, err)
		assert.NotNil(t, testSession)
		assert.NotZero(t, testSession.LastSeenAt)
		assert.Equal(t, expected["principal"], testSession.Principal)
		assert.Equal(t, expected["hostname"], testSession.Hostname)
		assert.Equal(t, expected["sessionIdentifier"], testSession.Identifier)
		assert.Equal(t, expected["hostIdentifier"], testSession.HostIdentifier)
		assert.Equal(t, expected["agentIdentifier"], testSession.AgentIdentifier)
	})

	/*
	 * Test when the `claimTasks` mutation is run by a session that already exists.
	 *
	 * Expected that the session is updated, and our test task is returned.
	 */
	t.Run("ExistingSession", func(t *testing.T) {
		expected := map[string]any{
			"principal":         "admin",
			"hostname":          "SOME_HOSTNAME",
			"sessionIdentifier": "SOME_ID",
			"hostIdentifier":    "SOME_HOST_ID",
			"agentIdentifier":   "SOME_AGENT_ID",
		}
		ids, err := claimTasks(expected)
		require.NoError(t, err)
		assert.Len(t, ids, 1)
		assert.Equal(t, testTasks[0].ID, ids[0])
		testTask := graph.Task.GetX(ctx, testTasks[0].ID)
		assert.NotZero(t, testTask.ClaimedAt)

		testSession, err := testTask.Session(ctx)
		require.NoError(t, err)
		assert.Equal(t, testSessions[0].ID, testSession.ID)
		assert.NotZero(t, testSession.LastSeenAt)
		assert.Equal(t, expected["principal"], testSession.Principal)
		assert.Equal(t, expected["hostname"], testSession.Hostname)
		assert.Equal(t, expected["sessionIdentifier"], testSession.Identifier)
		assert.Equal(t, expected["hostIdentifier"], testSession.HostIdentifier)
		assert.Equal(t, expected["agentIdentifier"], testSession.AgentIdentifier)
	})

}
