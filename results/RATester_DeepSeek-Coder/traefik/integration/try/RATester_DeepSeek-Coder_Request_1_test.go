package try

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestRequest_1(t *testing.T) {
	t.Parallel()

	type args struct {
		req        *http.Request
		timeout    time.Duration
		conditions []ResponseCondition
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				req:     &http.Request{},
				timeout: time.Duration(10),
				conditions: []ResponseCondition{
					func(resp *http.Response) error {
						if resp.StatusCode != 200 {
							return errors.New("unexpected status code")
						}
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				req:     &http.Request{},
				timeout: time.Duration(10),
				conditions: []ResponseCondition{
					func(resp *http.Response) error {
						if resp.StatusCode != 404 {
							return errors.New("unexpected status code")
						}
						return nil
					},
				},
			},
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

			err := Request(tt.args.req, tt.args.timeout, tt.args.conditions...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
