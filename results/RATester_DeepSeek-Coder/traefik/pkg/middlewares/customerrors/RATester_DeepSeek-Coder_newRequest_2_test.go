package customerrors

import (
	"fmt"
	"testing"
)

func TestNewRequest_2(t *testing.T) {
	type args struct {
		baseURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				baseURL: "http://example.com",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				baseURL: ":",
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

			_, err := newRequest(tt.args.baseURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("newRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
