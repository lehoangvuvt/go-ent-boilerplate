package bootstrapstack

import (
	mailports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/mail"
	queueports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/queue"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	httpuser "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/user"
	userusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/user"
)

type BuildUserStackArgs struct {
	UserRepository repositoryports.UserRepository

	MailService mailports.MailService

	QueueProducer queueports.QueueProducer
	QueueConsumer queueports.QueueConsumer
	QueueCloser   queueports.QueueCloser
	QueueHandler  queueports.QueueHandler
}

func BuildUserStack(args BuildUserStackArgs) *httpuser.UserHandler {
	createUserUC := userusecase.NewUserUsecase(args.UserRepository, args.MailService, args.QueueProducer)
	handler := httpuser.NewUserHandler(httpuser.NewUserHandlerArgs{
		CreateUserUC: createUserUC,
	})
	return handler
}
