package safe

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestNew_23(t *testing.T) {
	t.Run("TestNew", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		type args struct {
			value interface{}
		}
		tests := []struct {
			name string
			args args
			want *Safe
		}{
			{
				name: "TestNew_1",
				args: args{
					value: 1,
				},
				want: &Safe{
					value: 1,
					lock:  sync.RWMutex{},
				},
			},
			{
				name: "TestNew_2",
				args: args{
					value: "test",
				},
				want: &Safe{
					value: "test",
					lock:  sync.RWMutex{},
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

				if got := New(tt.args.value); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
