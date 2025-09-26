package try

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestBodyContainsOr_1(t *testing.T) {
	tests := []struct {
		name    string
		values  []string
		body    string
		wantErr bool
	}{
		{
			name:    "contains",
			values:  []string{"foo", "bar"},
			body:    "foo bar baz",
			wantErr: false,
		},
		{
			name:    "does not contain",
			values:  []string{"foo", "baz"},
			body:    "bar",
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

			err := BodyContainsOr(tt.values...)(res)
			if (err != nil) != tt.wantErr {
				t.Errorf("BodyContainsOr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
