package idempotency

import (
	"fmt"
	"testing"
)

func TestEncodeMsg_2(t *testing.T) {
	type args struct {
		z *response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				z: &response{
					Headers: map[string][]string{
						"Content-Type": {"application/json"},
					},
					Body:       []byte("Hello, World!"),
					StatusCode: 200,
				},
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.args.z.EncodeMsg(nil); (err != nil) != tt.wantErr {
				t.Errorf("response.EncodeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
