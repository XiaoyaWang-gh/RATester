package hcontext

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewContextDispatcher_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want ContextDispatcher[any]
	}{
		{
			name: "Test case 1",
			args: args{key: "test_key"},
			want: NewContextDispatcher[any]("test_key"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewContextDispatcher[any](tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContextDispatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
