package tcp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewChain_1(t *testing.T) {
	type args struct {
		constructors []Constructor
	}
	tests := []struct {
		name string
		args args
		want Chain
	}{
		{
			name: "Test Case 1",
			args: args{
				constructors: []Constructor{
					func(h Handler) (Handler, error) {
						return h, nil
					},
				},
			},
			want: Chain{
				constructors: []Constructor{
					func(h Handler) (Handler, error) {
						return h, nil
					},
				},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewChain(tt.args.constructors...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
