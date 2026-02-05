package pubsub

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/protobuf/proto"
	"realm.pub/tavern/portals/portalpb"
)

type memDriver struct {
	mu     sync.RWMutex
	topics map[string]*memTopic
}

type memTopic struct {
	mu          sync.RWMutex
	subscribers []*memSubscription
}

type memSubscription struct {
	ch chan *portalpb.Mote
}

// EnsurePublisher creates and returns an in-memory Publisher for the specified topic.
func (drv *memDriver) EnsurePublisher(ctx context.Context, topic string) (Publisher, error) {
	drv.mu.Lock()
	defer drv.mu.Unlock()

	if drv.topics == nil {
		drv.topics = make(map[string]*memTopic)
	}

	t, ok := drv.topics[topic]
	if !ok {
		t = &memTopic{}
		drv.topics[topic] = t
	}

	return &memPublisher{topic: t}, nil
}

// EnsureSubscriber creates and returns an in-memory Subscriber for the specified topic and subscription.
func (drv *memDriver) EnsureSubscriber(ctx context.Context, topic, subscription string) (Subscriber, error) {
	drv.mu.Lock()
	defer drv.mu.Unlock()

	if drv.topics == nil {
		drv.topics = make(map[string]*memTopic)
	}

	t, ok := drv.topics[topic]
	if !ok {
		t = &memTopic{}
		drv.topics[topic] = t
	}

	sub := &memSubscription{
		ch: make(chan *portalpb.Mote, 256),
	}
	t.mu.Lock()
	t.subscribers = append(t.subscribers, sub)
	t.mu.Unlock()

	return sub, nil
}

// Close is a no-op for the in-memory driver.
func (drv *memDriver) Close() error {
	return nil
}

type memPublisher struct {
	topic *memTopic
}

func (pub *memPublisher) Publish(ctx context.Context, mote *portalpb.Mote) error {
	// Deep copy the mote to prevent shared state issues
	data, err := proto.Marshal(mote)
	if err != nil {
		return fmt.Errorf("failed to marshal mote: %w", err)
	}

	pub.topic.mu.RLock()
	defer pub.topic.mu.RUnlock()

	for _, sub := range pub.topic.subscribers {
		clone := &portalpb.Mote{}
		if err := proto.Unmarshal(data, clone); err != nil {
			return fmt.Errorf("failed to unmarshal mote: %w", err)
		}

		select {
		case sub.ch <- clone:
		default:
			// Drop message if subscriber buffer is full
		}
	}

	return nil
}

func (sub *memSubscription) Receive(ctx context.Context, f func(context.Context, *portalpb.Mote)) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case mote := <-sub.ch:
			f(ctx, mote)
		}
	}
}
