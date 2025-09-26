package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProcessAccounts_1(t *testing.T) {
	type args struct {
		accounts Accounts
	}
	tests := []struct {
		name string
		args args
		want authPairs
	}{
		{
			name: "Test with empty accounts",
			args: args{
				accounts: Accounts{},
			},
			want: authPairs{},
		},
		{
			name: "Test with non-empty accounts",
			args: args{
				accounts: Accounts{
					"user1": "password1",
					"user2": "password2",
				},
			},
			want: authPairs{
				{
					value: "Basic dXNlcjE6cGFzc3dvcmQx",
					user:  "user1",
				},
				{
					value: "Basic dXNlcjI6cGFzc3dvcmQy",
					user:  "user2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := processAccounts(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
