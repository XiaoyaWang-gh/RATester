package safe

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestNew_23(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want *Safe
	}{
		{
			name: "Test New with int",
			args: args{value: 1},
			want: &Safe{value: 1, lock: sync.RWMutex{}},
		},
		{
			name: "Test New with string",
			args: args{value: "test"},
			want: &Safe{value: "test", lock: sync.RWMutex{}},
		},
		{
			name: "Test New with nil",
			args: args{value: nil},
			want: &Safe{value: nil, lock: sync.RWMutex{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := New(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
