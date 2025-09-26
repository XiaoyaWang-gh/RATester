package conn

import (
	"fmt"
	"testing"
)

func TestSetAlive_1(t *testing.T) {
	type args struct {
		tp string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				tp: "test1",
			},
		},
		{
			name: "test2",
			args: args{
				tp: "test2",
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

			s := &Conn{}
			s.SetAlive(tt.args.tp)
		})
	}
}
