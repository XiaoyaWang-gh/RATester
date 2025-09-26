package try

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStatusCodeIs_1(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		response *http.Response
		wantErr  bool
	}{
		{
			name:     "Status code is 200",
			status:   200,
			response: &http.Response{StatusCode: 200},
			wantErr:  false,
		},
		{
			name:     "Status code is not 200",
			status:   200,
			response: &http.Response{StatusCode: 404},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			condition := StatusCodeIs(tt.status)
			err := condition(tt.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatusCodeIs() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
