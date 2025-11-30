package rabbitmq

import (
	"context"
	"errors"
	"log"
	"time"

	queueports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(dsn string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
	}, nil
}

func (r *RabbitMQ) Close() error {
	if r.channel != nil {
		_ = r.channel.Close()
	}
	if r.conn != nil {
		return r.conn.Close()
	}
	return nil
}

func (r *RabbitMQ) Publish(ctx context.Context, topic string, payload []byte) error {
	return r.channel.PublishWithContext(
		ctx,
		"amq.direct",
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
			Timestamp:   time.Now(),
		},
	)
}

func (r *RabbitMQ) Consume(ctx context.Context, queue string, handler queueports.QueueHandler) error {
	_, err := r.channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := r.channel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			log.Printf("[RabbitMQ] Consumer stopped (%s)\n", queue)
			return nil

		case msg, ok := <-msgs:
			if !ok {
				log.Println("[RabbitMQ] Channel closed")
				return errors.New("channel closed")
			}

			go func(m amqp.Delivery) {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("[RabbitMQ] panic recovered: %v", r)
						_ = m.Nack(false, true)
					}
				}()

				qmsg := queueports.QueueMessage{
					ID:   m.MessageId,
					Body: m.Body,
					Headers: map[string]string{
						"content_type": m.ContentType,
					},
				}

				err := handler.HandleMessage(ctx, qmsg)
				if err != nil {
					_ = m.Nack(false, true)
					return
				}

				_ = m.Ack(false)
			}(msg)
		}
	}
}
