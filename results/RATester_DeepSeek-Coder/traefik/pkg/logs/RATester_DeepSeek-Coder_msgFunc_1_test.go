package logs

import (
	"fmt"
	"testing"
)

func TestMsgFunc_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test case 1",
			want: "test1",
		},
		{
			name: "Test case 2",
			want: "test2",
		},
		{
			name: "Test case 3",
			want: "test3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := msgFunc(tt.want); got() != tt.want {
				t.Errorf("msgFunc() = %v, want %v", got(), tt.want)
			}
		})
	}
}
