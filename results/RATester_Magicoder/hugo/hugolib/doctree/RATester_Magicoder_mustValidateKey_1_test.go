package doctree

import (
	"fmt"
	"testing"
)

func TestmustValidateKey_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				key: "valid_key",
			},
			want: "valid_key",
		},
		{
			name: "Test Case 2",
			args: args{
				key: "invalid_key",
			},
			want: "invalid_key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := mustValidateKey(tt.args.key); got != tt.want {
				t.Errorf("mustValidateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
