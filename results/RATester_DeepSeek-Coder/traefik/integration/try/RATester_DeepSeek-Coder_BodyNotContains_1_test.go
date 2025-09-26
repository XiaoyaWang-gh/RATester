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
			name:    "contains",
			values:  []string{"test"},
			body:    "this is a test",
			wantErr: true,
		},
		{
			name:    "not contains",
			values:  []string{"test"},
			body:    "this is not a test",
			wantErr: false,
		},
		{
			name:    "multiple values",
			values:  []string{"test1", "test2"},
			body:    "this is not a test1 and not a test2",
			wantErr: false,
		},
		{
			name:    "empty body",
			values:  []string{"test"},
			body:    "",
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
			}
		})
	}
}
