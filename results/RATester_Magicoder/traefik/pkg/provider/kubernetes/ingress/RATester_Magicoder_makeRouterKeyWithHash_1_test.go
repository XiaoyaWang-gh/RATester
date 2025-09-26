package ingress

import (
	"fmt"
	"testing"
)

func TestMakeRouterKeyWithHash_1(t *testing.T) {
	type args struct {
		key  string
		rule string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				key:  "test_key",
				rule: "test_rule",
			},
			want:    "test_key-746861745f72756c65",
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				key:  "another_key",
				rule: "another_rule",
			},
			want:    "another_key-616e6f746865725f72756c65",
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

			got, err := makeRouterKeyWithHash(tt.args.key, tt.args.rule)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeRouterKeyWithHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("makeRouterKeyWithHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
