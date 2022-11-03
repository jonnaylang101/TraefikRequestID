package traefikrequestid

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

const reqIdHeaderKey string = "X-Request-ID"

type Config struct{}

func CreateConfig() *Config {
	return &Config{}
}

type RequestID struct {
	next http.Handler
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &RequestID{
		next: next,
		name: name,
	}, nil
}

func (rid *RequestID) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello world!\n"))
	// rid.next.ServeHTTP(rw, req)
}

func addRequestIdToHeaders(httpRequest *http.Request, requestId uuid.UUID) error {
	if httpRequest == nil {
		panic("addIdToHeaders: provided httpRequest param has nil value")
	}
	if requestId == uuid.Nil {
		return errors.New("addIdToHeaders: provided requestId param has nil value")
	}

	httpRequest.Header.Set(reqIdHeaderKey, requestId.String())

	return nil
}
