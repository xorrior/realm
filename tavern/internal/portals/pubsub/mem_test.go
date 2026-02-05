package pubsub_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"realm.pub/tavern/internal/portals/pubsub"
	"realm.pub/tavern/portals/portalpb"
)

func TestInMemoryDriver(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := pubsub.NewClient()
	defer c.Close()

	// Test EnsurePublisher
	publisher, err := c.EnsurePublisher(ctx, "test-topic")
	require.NoError(t, err)
	require.NotNil(t, publisher)

	// Test EnsureSubscriber
	subscriber, err := c.EnsureSubscriber(ctx, "test-topic", "test-sub")
	require.NoError(t, err)
	require.NotNil(t, subscriber)

	// Test Publish and Receive
	received := make(chan *portalpb.Mote, 1)

	go func() {
		err := subscriber.Receive(ctx, func(ctx context.Context, mote *portalpb.Mote) {
			received <- mote
		})
		if err != nil && ctx.Err() == nil {
			t.Logf("Receive stopped: %v", err)
		}
	}()

	mote := &portalpb.Mote{
		Payload: &portalpb.Mote_Bytes{
			Bytes: &portalpb.BytesPayload{
				Kind: portalpb.BytesPayloadKind_BYTES_PAYLOAD_KIND_DATA,
				Data: []byte("hello world"),
			},
		},
	}

	err = publisher.Publish(ctx, mote)
	require.NoError(t, err)

	select {
	case r := <-received:
		assert.Equal(t, portalpb.BytesPayloadKind_BYTES_PAYLOAD_KIND_DATA, r.GetBytes().Kind)
		assert.Equal(t, []byte("hello world"), r.GetBytes().Data)
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for message")
	}
}

func TestInMemoryDriver_MultipleSubscribers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := pubsub.NewClient()
	defer c.Close()

	publisher, err := c.EnsurePublisher(ctx, "multi-topic")
	require.NoError(t, err)

	sub1, err := c.EnsureSubscriber(ctx, "multi-topic", "sub-1")
	require.NoError(t, err)
	sub2, err := c.EnsureSubscriber(ctx, "multi-topic", "sub-2")
	require.NoError(t, err)

	received1 := make(chan *portalpb.Mote, 1)
	received2 := make(chan *portalpb.Mote, 1)

	go sub1.Receive(ctx, func(ctx context.Context, mote *portalpb.Mote) {
		received1 <- mote
	})
	go sub2.Receive(ctx, func(ctx context.Context, mote *portalpb.Mote) {
		received2 <- mote
	})

	// Small delay to ensure receivers are ready
	time.Sleep(10 * time.Millisecond)

	mote := &portalpb.Mote{
		Payload: &portalpb.Mote_Bytes{
			Bytes: &portalpb.BytesPayload{
				Kind: portalpb.BytesPayloadKind_BYTES_PAYLOAD_KIND_DATA,
				Data: []byte("broadcast"),
			},
		},
	}

	err = publisher.Publish(ctx, mote)
	require.NoError(t, err)

	for _, ch := range []chan *portalpb.Mote{received1, received2} {
		select {
		case r := <-ch:
			assert.Equal(t, []byte("broadcast"), r.GetBytes().Data)
		case <-time.After(5 * time.Second):
			t.Fatal("timed out waiting for message")
		}
	}
}
