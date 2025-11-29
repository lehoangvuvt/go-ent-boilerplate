package httpmiddleware

import (
	"bytes"
	"crypto/sha256"
	"io"
	"net/http"

	idempotencyports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/idempotency"
)

type IdempotencyMiddleware struct {
	store idempotencyports.IdempotencyStore
}

func NewIdempotencyMiddleware(store idempotencyports.IdempotencyStore) *IdempotencyMiddleware {
	return &IdempotencyMiddleware{store: store}
}

func hashRequestBody(body []byte) string {
	h := sha256.Sum256(body)
	return string(h[:])
}

func (m *IdempotencyMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Idempotency-Key")

		if key == "" {
			next.ServeHTTP(w, r)
			return
		}

		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		reqHash := hashRequestBody(bodyBytes)

		rec, _ := m.store.Get(r.Context(), key)
		if rec != nil {
			switch rec.Status {
			case "pending":
				http.Error(w, "Request is already processing", http.StatusConflict)
				return

			case "success":
				w.Header().Set("Content-Type", "application/json")
				w.Write(rec.Response)
				return
			}
		}

		_ = m.store.SavePending(r.Context(), key, reqHash)

		recorder := NewResponseRecorder(w)
		next.ServeHTTP(recorder, r)

		if recorder.Status >= 200 && recorder.Status < 300 {
			_ = m.store.SaveSuccess(r.Context(), key, recorder.Body.Bytes())
		} else {
			_ = m.store.SaveFailed(r.Context(), key)
		}
	})
}
