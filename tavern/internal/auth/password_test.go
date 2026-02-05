package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"realm.pub/tavern/internal/auth"
	"realm.pub/tavern/internal/ent/enttest"
)

func TestSetupHandler_FirstUser(t *testing.T) {
	graph := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer graph.Close()

	handler := auth.NewSetupHandler(graph)

	// GET should indicate setup is needed
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, httptest.NewRequest(http.MethodGet, "/auth/setup", nil))
	result := resp.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)
	var getResp map[string]interface{}
	json.NewDecoder(result.Body).Decode(&getResp)
	assert.True(t, getResp["needsSetup"].(bool))

	// POST to create first user
	body, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": "securepassword123",
	})
	resp = httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/setup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(resp, req)
	result = resp.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	// Verify user was created
	assert.Equal(t, 1, graph.User.Query().CountX(context.Background()))

	// Verify session cookie was set
	var sessionCookie *http.Cookie
	for _, cookie := range result.Cookies() {
		if cookie.Name == auth.SessionCookieName {
			sessionCookie = cookie
			break
		}
	}
	require.NotNil(t, sessionCookie)
	assert.NotEmpty(t, sessionCookie.Value)

	// POST again should fail (setup only allowed once)
	body, _ = json.Marshal(map[string]string{
		"username": "hacker",
		"password": "tryingtohack123",
	})
	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/auth/setup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(resp, req)
	result = resp.Result()
	assert.Equal(t, http.StatusForbidden, result.StatusCode)
}

func TestPasswordLoginHandler(t *testing.T) {
	graph := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer graph.Close()

	// Create a user with a password
	hash, err := auth.HashPassword("testpassword123")
	require.NoError(t, err)
	graph.User.Create().
		SetName("testuser").
		SetPasswordHash(hash).
		SetPhotoURL("").
		SetIsAdmin(true).
		SetIsActivated(true).
		SaveX(context.Background())

	handler := auth.NewPasswordLoginHandler(graph)

	// Test successful login
	body, _ := json.Marshal(map[string]string{
		"username": "testuser",
		"password": "testpassword123",
	})
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(resp, req)
	result := resp.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	// Verify session cookie
	var sessionCookie *http.Cookie
	for _, cookie := range result.Cookies() {
		if cookie.Name == auth.SessionCookieName {
			sessionCookie = cookie
			break
		}
	}
	require.NotNil(t, sessionCookie)

	// Test wrong password
	body, _ = json.Marshal(map[string]string{
		"username": "testuser",
		"password": "wrongpassword",
	})
	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(resp, req)
	result = resp.Result()
	assert.Equal(t, http.StatusUnauthorized, result.StatusCode)

	// Test nonexistent user
	body, _ = json.Marshal(map[string]string{
		"username": "nonexistent",
		"password": "somepassword",
	})
	resp = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(resp, req)
	result = resp.Result()
	assert.Equal(t, http.StatusUnauthorized, result.StatusCode)
}

func TestHashPassword_And_CheckPassword(t *testing.T) {
	password := "mysecurepassword"
	hash, err := auth.HashPassword(password)
	require.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)

	assert.True(t, auth.CheckPassword(hash, password))
	assert.False(t, auth.CheckPassword(hash, "wrongpassword"))
}

func TestSetupHandler_PasswordTooShort(t *testing.T) {
	graph := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer graph.Close()

	handler := auth.NewSetupHandler(graph)

	body, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": "short",
	})
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/setup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(resp, req)
	result := resp.Result()
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)
}
