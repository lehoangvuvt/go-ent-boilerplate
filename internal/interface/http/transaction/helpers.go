package httptransaction

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	transactionmethoddomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction/method"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
)

func parseTransactionFilter(r *http.Request) (repositoryports.TransactionFilter, error) {
	q := r.URL.Query()

	var filter repositoryports.TransactionFilter

	if v := q.Get("user_id"); v != "" {
		uid, err := uuid.Parse(v)
		if err != nil {
			return filter, errors.New("invalid user_id")
		}
		filter.UserID = &uid
	}

	if v := q.Get("status"); v != "" {
		status := transactiondomain.TransactionStatus(v)
		switch status {
		case transactiondomain.Pending,
			transactiondomain.Completed,
			transactiondomain.Failed,
			transactiondomain.Rejected:
			filter.Status = &status
		default:
			return filter, errors.New("invalid status")
		}
	}

	if v := q.Get("method"); v != "" {
		m := transactionmethoddomain.PaymentMethodType(v)
		switch m {
		case transactionmethoddomain.MethodVisa,
			transactionmethoddomain.MethodBanking,
			transactionmethoddomain.MethodEWallet,
			transactionmethoddomain.MethodQR:
			filter.Method = &m
		default:
			return filter, errors.New("invalid method")
		}
	}

	if v := q.Get("limit"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 0 {
			return filter, errors.New("invalid limit")
		}
		filter.Limit = n
	}

	if v := q.Get("offset"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 0 {
			return filter, errors.New("invalid offset")
		}
		filter.Offset = n
	}

	return filter, nil
}
