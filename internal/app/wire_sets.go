package app

import (
	"github.com/google/wire"
	"github.com/lehoangvuvt/go-ent-boilerplate/internal/bootstrap"
)

var DBSet = wire.NewSet(
	provideEntDB,
	provideEntClient,
)

var RepositorySet = wire.NewSet(
	provideUserRepository,
	provideTransactionRepository,
	wire.Struct(new(bootstrap.Repositories), "*"),
)

var ServiceSet = wire.NewSet(
	provideJWTDuration,
	provideJWTService,

	provideRedisOptions,
	provideCache,

	provideResendMailService,

	provideQueueAdapter,
	provideQueueProducer,
	provideQueueConsumer,
	provideQueueCloser,
	provideRegisterEmailHandler,
	provideQueueConsumerRunner,

	wire.Struct(new(bootstrap.Services), "*"),
)

var StoreSet = wire.NewSet(
	provideIdempotencyStore,
	wire.Struct(new(bootstrap.Stores), "*"),
)

var HandlerSet = wire.NewSet(
	provideHandlerArgs,
	bootstrap.BootstrapHandler,
)

var ContainerSet = wire.NewSet(
	wire.Struct(new(Container), "*"),
)

var WorkerSet = wire.NewSet(
	provideQueueAdapter,
	provideQueueConsumer,
	provideQueueCloser,
	provideResendMailService,
	provideRegisterEmailHandler,
	provideWorkerRunner,
)
