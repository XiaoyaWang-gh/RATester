package try

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestBodyNotContains_1(t *testing.T) {
	tests := []struct {
		name    string
		values  []string
		body    string
		wantErr bool
	}{
		{
			name:    "Body contains value",
			values:  []string{"test"},
			body:    "This is a test body",
			wantErr: true,
		},
		{
			name:    "Body does not contain value",
			values:  []string{"test"},
			body:    "This is a body",
			wantErr: false,
		},
		{
			name:    "Body contains multiple values",
			values:  []string{"test", "body"},
			body:    "This is a test body",
			wantErr: true,
		},
		{
			name:    "Body does not contain any value",
			values:  []string{"test", "body"},
			body:    "This is a different body",
			wantErr: false,
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

			err := BodyNotContains(tt.values...)(res)
			if (err != nil) != tt.wantErr {
				t.Errorf("BodyNotContains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
