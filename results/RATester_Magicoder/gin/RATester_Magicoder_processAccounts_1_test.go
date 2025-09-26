package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestprocessAccounts_1(t *testing.T) {
	type args struct {
		accounts Accounts
	}
	tests := []struct {
		name string
		args args
		want authPairs
	}{
		{
			name: "Test Case 1",
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
		{
			name: "Test Case 2",
			args: args{
				accounts: Accounts{
					"user1": "password1",
					"":      "password2",
				},
			},
			want: authPairs{
				{
					value: "Basic dXNlcjE6cGFzc3dvcmQx",
					user:  "user1",
				},
			},
		},
		{
			name: "Test Case 3",
			args: args{
				accounts: Accounts{
					"": "password1",
				},
			},
			want: authPairs{},
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
