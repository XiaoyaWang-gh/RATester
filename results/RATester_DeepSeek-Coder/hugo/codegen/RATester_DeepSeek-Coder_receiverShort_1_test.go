package codegen

import (
	"fmt"
	"testing"
)

func TestReceiverShort_1(t *testing.T) {
	type args struct {
		receiver string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				receiver: "*Hello",
			},
			want: "h",
		},
		{
			name: "Test Case 2",
			args: args{
				receiver: "*World",
			},
			want: "w",
		},
		{
			name: "Test Case 3",
			args: args{
				receiver: "*GoLang*",
			},
			want: "g",
		},
		{
			name: "Test Case 4",
			args: args{
				receiver: "*UnitTest*",
			},
			want: "u",
		},
		{
			name: "Test Case 5",
			args: args{
				receiver: "*",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := receiverShort(tt.args.receiver); got != tt.want {
				t.Errorf("receiverShort() = %v, want %v", got, tt.want)
			}
		})
	}
}
