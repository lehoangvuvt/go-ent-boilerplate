package bootstrapstack

import (
	"context"

	queueports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/queue"
)

type WorkerRunner func(ctx context.Context) error

type BuildWorkerStackArgs struct {
	QueueConsumer queueports.QueueConsumer
	QueueCloser   queueports.QueueCloser
	QueueHandler  queueports.QueueHandler
}

func BuildWorkerStack(args BuildWorkerStackArgs) WorkerRunner {
	return func(ctx context.Context) error {
		return args.QueueConsumer.Consume(
			ctx,
			"send_email",
			args.QueueHandler,
			&queueports.ConsumeOptions{
				AutoAck:     false,
				Concurrency: 5,
				Prefetch:    5,
			},
		)
	}
}
