package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCustomRecovery_1(t *testing.T) {
	type args struct {
		handle RecoveryFunc
	}
	tests := []struct {
		name string
		args args
		want HandlerFunc
	}{
		{
			name: "Test case 1",
			args: args{
				handle: func(c *Context, err any) {
					// Test logic here
				},
			},
			want: func(c *Context) {
				// Test logic here
			},
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CustomRecovery(tt.args.handle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomRecovery() = %v, want %v", got, tt.want)
			}
		})
	}
}
