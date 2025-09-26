package common

import (
	"fmt"
	"testing"
)

func TestGetIntNoErrByStr_1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				str: "123",
			},
			want: 123,
		},
		{
			name: "Test case 2",
			args: args{
				str: "456",
			},
			want: 456,
		},
		{
			name: "Test case 3",
			args: args{
				str: "789",
			},
			want: 789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetIntNoErrByStr(tt.args.str); got != tt.want {
				t.Errorf("GetIntNoErrByStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
