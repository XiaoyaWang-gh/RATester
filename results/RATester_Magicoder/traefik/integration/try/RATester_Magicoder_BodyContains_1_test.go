package try

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestBodyContains_1(t *testing.T) {
	tests := []struct {
		name    string
		values  []string
		body    string
		wantErr bool
	}{
		{
			name:    "success",
			values:  []string{"test"},
			body:    "this is a test body",
			wantErr: false,
		},
		{
			name:    "failure",
			values:  []string{"fail"},
			body:    "this is a test body",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			res := &http.Response{
				Body: io.NopCloser(strings.NewReader(tt.body)),
			}

			err := BodyContains(tt.values...)(res)
			if (err != nil) != tt.wantErr {
				t.Errorf("BodyContains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
