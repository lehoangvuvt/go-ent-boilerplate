package queueports

import "context"

type QueueMessage struct {
	ID      string
	Body    []byte
	Headers map[string]string
}

type QueueProducer interface {
	Publish(ctx context.Context, topic string, payload []byte) error
}

type QueueConsumer interface {
	Consume(ctx context.Context, topic string, handler QueueHandler) error
}

type QueueHandler interface {
	HandleMessage(ctx context.Context, msg QueueMessage) error
}

type QueueCloser interface {
	Close() error
}
