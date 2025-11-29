package idempotencyports

import "context"

type IdempotencyRecord struct {
	Key         string
	RequestHash string
	Response    []byte
	Status      string
}

type IdempotencyStore interface {
	Get(ctx context.Context, key string) (*IdempotencyRecord, error)
	SavePending(ctx context.Context, key string, requestHash string) error
	SaveSuccess(ctx context.Context, key string, response []byte) error
	SaveFailed(ctx context.Context, key string) error
}
