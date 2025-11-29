package httpmiddleware

import (
	"bytes"
	"net/http"
)

type ResponseRecorder struct {
	http.ResponseWriter
	Body   *bytes.Buffer
	Status int
}

func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{
		ResponseWriter: w,
		Body:           &bytes.Buffer{},
		Status:         http.StatusOK,
	}
}

func (rr *ResponseRecorder) WriteHeader(status int) {
	rr.Status = status
	rr.ResponseWriter.WriteHeader(status)
}

func (rr *ResponseRecorder) Write(b []byte) (int, error) {
	rr.Body.Write(b)
	return rr.ResponseWriter.Write(b)
}
