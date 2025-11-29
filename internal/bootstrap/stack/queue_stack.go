package bootstrapstack

import (
	"log"

	"github.com/lehoangvuvt/go-ent-boilerplate/internal/infrastructure/queue/rabbitmq"
	queueports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/queue"
)

type BuildQueueStackArgs struct {
	DSN string
}

func BuildQueueStack(args BuildQueueStackArgs) (queueports.QueueProducer, queueports.QueueConsumer, queueports.QueueCloser) {
	q, err := rabbitmq.NewRabbitMQ(args.DSN)
	if err != nil {
		log.Fatalf("RabbitMQ connection failed: %v", err)
	}
	return q, q, q
}
