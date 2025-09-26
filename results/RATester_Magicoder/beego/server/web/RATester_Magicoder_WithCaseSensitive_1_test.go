package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithCaseSensitive_1(t *testing.T) {
	type args struct {
		sensitive bool
	}
	tests := []struct {
		name string
		args args
		want FilterOpt
	}{
		{
			name: "Test case 1",
			args: args{
				sensitive: true,
			},
			want: WithCaseSensitive(true),
		},
		{
			name: "Test case 2",
			args: args{
				sensitive: false,
			},
			want: WithCaseSensitive(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WithCaseSensitive(tt.args.sensitive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCaseSensitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
