package traefikrequestid

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

const reqIdHeaderKey string = "X-Request-ID"

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
