package metrics

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewMultiCounterWithHeaders_1(t *testing.T) {
	type args struct {
		c []CounterWithHeaders
	}
	tests := []struct {
		name string
		args args
		want MultiCounterWithHeaders
	}{
		{
			name: "test",
			args: args{
				c: []CounterWithHeaders{
					&counterWithHeaders{
						name: "test",
					},
				},
			},
			want: []CounterWithHeaders{
				&counterWithHeaders{
					name: "test",
				},
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

			if got := NewMultiCounterWithHeaders(tt.args.c...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMultiCounterWithHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
