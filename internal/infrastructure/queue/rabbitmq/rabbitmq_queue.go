package rabbitmqqueue

import (
	"context"
	"fmt"

	queueports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/queue"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQQueue struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

func NewRabbitMQQueue(dsn string) (*RabbitMQQueue, error) {
	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq connect error: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("rabbitmq open channel error: %w", err)
	}

	return &RabbitMQQueue{
		conn:    conn,
		channel: ch,
	}, nil
}

var _ queueports.QueueProducer = (*RabbitMQQueue)(nil)

func (r *RabbitMQQueue) Publish(
	ctx context.Context,
	topic string,
	payload []byte,
	opts *queueports.PublishOptions,
) error {
	if _, err := r.channel.QueueDeclare(
		topic,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("rabbitmq queue declare error: %w", err)
	}

	headers := amqp091.Table{}
	if opts != nil && opts.Headers != nil {
		for k, v := range opts.Headers {
			headers[k] = v
		}
	}

	return r.channel.PublishWithContext(
		ctx,
		"",
		topic,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        payload,
			Headers:     headers,
		},
	)
}

var _ queueports.QueueConsumer = (*RabbitMQQueue)(nil)

func (r *RabbitMQQueue) Consume(
	ctx context.Context,
	topic string,
	handler queueports.QueueHandler,
	opts *queueports.ConsumeOptions,
) error {
	if opts == nil {
		opts = &queueports.ConsumeOptions{}
	}

	if _, err := r.channel.QueueDeclare(
		topic,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("rabbitmq queue declare error: %w", err)
	}

	// QoS / prefetch
	if opts.Prefetch > 0 {
		if err := r.channel.Qos(opts.Prefetch, 0, false); err != nil {
			return fmt.Errorf("rabbitmq qos error: %w", err)
		}
	}

	msgs, err := r.channel.Consume(
		topic,
		"",
		opts.AutoAck,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("rabbitmq consume error: %w", err)
	}
	workerCount := opts.Concurrency
	if workerCount <= 0 {
		workerCount = 1
	}

	errChan := make(chan error, 1)

	for i := 0; i < workerCount; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case d, ok := <-msgs:
					if !ok {
						return
					}

					msg := queueports.QueueMessage{
						ID:   d.MessageId,
						Body: d.Body,
						Headers: func() map[string]string {
							h := map[string]string{}
							for k, v := range d.Headers {
								if s, ok := v.(string); ok {
									h[k] = s
								}
							}
							return h
						}(),
					}

					handleErr := handler.HandleMessage(ctx, msg)

					if opts.AutoAck {
						continue
					}

					if handleErr != nil {
						_ = d.Nack(false, true)

						select {
						case errChan <- handleErr:
						default:
						}
						continue
					}

					_ = d.Ack(false)
				}
			}
		}()
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return fmt.Errorf("handler error: %w", err)
	}
}

var _ queueports.QueueCloser = (*RabbitMQQueue)(nil)

func (r *RabbitMQQueue) Close() error {
	var err error

	if r.channel != nil {
		if e := r.channel.Close(); e != nil {
			err = e
		}
	}

	if r.conn != nil {
		if e := r.conn.Close(); e != nil && err == nil {
			err = e
		}
	}

	return err
}
