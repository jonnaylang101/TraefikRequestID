package traefikrequestid

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func Test_addRequestIdToHeaders(t *testing.T) {
	newUUID := uuid.New()

	type args struct {
		req *http.Request
		id  uuid.UUID
	}
	tests := []struct {
		when      string
		it        string
		args      args
		wantErr   bool
		wantPanic bool
		wantReqId string
	}{
		{
			when: "the http request has a nil value",
			it:   "should panic",
			args: args{
				req: nil,
				id:  uuid.New(),
			},
			wantErr:   false,
			wantPanic: true,
		},
		{
			when: "the request id has nil value",
			it:   "should error and not insert the request id",
			args: args{
				req: httptest.NewRequest("GET", "/", nil),
				id:  uuid.Nil,
			},
			wantErr: true,
		},
		{
			when: "the request is well formed",
			it:   "should add the requestId header to the request",
			args: args{
				req: httptest.NewRequest("GET", "/", nil),
				id:  newUUID,
			},
			wantErr:   false,
			wantReqId: newUUID.String(),
		},
	}
	for _, tt := range tests {
		t.Run("When "+tt.when, func(t *testing.T) {
			t.Logf("\tIt %s", tt.it)
			defer func() {
				if panErr := recover(); (panErr != nil) != tt.wantPanic {
					t.Error("expected a panic but got none")
				}
			}()

			if err := addRequestIdToHeaders(tt.args.req, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("addIdToHeaders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotReqId := tt.args.req.Header.Get(reqIdHeaderKey)
			if gotReqId != tt.wantReqId {
				t.Errorf("wanted requestId to be %s but got %s", tt.wantReqId, gotReqId)
			}
		})
	}
}
