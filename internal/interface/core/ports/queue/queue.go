package queueports

import "context"

type QueueMessage struct {
	ID      string
	Body    []byte
	Headers map[string]string
}

type PublishOptions struct {
	Headers map[string]string
}

type ConsumeOptions struct {
	GroupID     string
	Concurrency int
	Prefetch    int
	AutoAck     bool
}
type QueueProducer interface {
	Publish(ctx context.Context, topic string, payload []byte, opts *PublishOptions) error
}

type QueueHandler interface {
	HandleMessage(ctx context.Context, msg QueueMessage) error
}
type QueueConsumer interface {
	Consume(ctx context.Context, topic string, handler QueueHandler, opts *ConsumeOptions) error
}

type QueueCloser interface {
	Close() error
}
