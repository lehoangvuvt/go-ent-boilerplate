package bootstrapstack

import (
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	httptransaction "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/transaction"
	transactionusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/transaction"
)

type BuildTransactionStackArgs struct {
	TransactionRepository repositoryports.TransactionRepository
}

func BuildTransactionStack(args BuildTransactionStackArgs) *httptransaction.TransactionHandler {
	createUC := transactionusecase.NewCreateTransactionUsecase(args.TransactionRepository)
	confirmUC := transactionusecase.NewConfirmTransactionUsecase(args.TransactionRepository)
	cancelUC := transactionusecase.NewCancelTransactionUsecase(args.TransactionRepository)
	failUC := transactionusecase.NewFailTransactionUsecase(args.TransactionRepository)
	findUC := transactionusecase.NewFindTransactionByIDUsecase(args.TransactionRepository)
	listUC := transactionusecase.NewListTransactionsUsecase(args.TransactionRepository)

	handler := httptransaction.NewTransactionHandler(httptransaction.NewTransactionHandlerArgs{
		CreateUC:  createUC,
		ConfirmUC: confirmUC,
		CancelUC:  cancelUC,
		FailUC:    failUC,
		FindUC:    findUC,
		ListUC:    listUC,
	})

	return handler
}
