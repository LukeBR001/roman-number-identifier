package pkg

import (
	"bytes"
	"encoding/json"
	"github.com/kinbiko/jsonassert"
	"identifier/pkg/model"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRomanIdentify(t *testing.T) {
	type args struct {
		payload model.TextPayload
	}

	ja := jsonassert.New(t)
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantResponse   string
	}{
		{
			name: "should return bigger roman number when successful",
			args: args{
				payload: model.TextPayload{
					Text: "XIVABIIIIQEMMDCCXXIVHGFLLXPOQLXXXQWEIXLQWEMmDcCcxXIvQWEICQWEXXL",
				},
			},
			wantStatusCode: http.StatusOK,
			wantResponse:   "2824",
		},
		{
			name: "should return bad request when not found any roman combination",
			args: args{
				payload: model.TextPayload{
					Text: "XIVIQWEVCQWEMMDCCILQWELXXXXQWEXLLQWEMmDcCcxXIvX",
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   `"not found valid roman combinations"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			payload, _ := json.Marshal(tt.args.payload)

			request := httptest.NewRequest("GET", "/text", bytes.NewReader(payload))

			RomanIdentify(rw, request)

			response := rw.Result()
			body, _ := io.ReadAll(response.Body)
			code := rw.Code

			if !reflect.DeepEqual(code, tt.wantStatusCode) {
				t.Errorf("handlePostInvoices() = %v, want %v", code, tt.wantStatusCode)
			}
			ja.Assertf(string(body), tt.wantResponse)

		})
	}
}
