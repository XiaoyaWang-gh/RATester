package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestchooseData_1(t *testing.T) {
	type args struct {
		custom   any
		wildcard any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test case 1: Both custom and wildcard are nil",
			args: args{
				custom:   nil,
				wildcard: nil,
			},
			want: nil,
		},
		{
			name: "Test case 2: Custom is not nil",
			args: args{
				custom:   "custom",
				wildcard: nil,
			},
			want: "custom",
		},
		{
			name: "Test case 3: Wildcard is not nil",
			args: args{
				custom:   nil,
				wildcard: "wildcard",
			},
			want: "wildcard",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := chooseData(tt.args.custom, tt.args.wildcard); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chooseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
