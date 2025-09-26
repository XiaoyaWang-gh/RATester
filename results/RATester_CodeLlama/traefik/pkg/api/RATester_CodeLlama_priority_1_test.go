package api

import (
	"fmt"
	"testing"
)

func TestPriority_1(t *testing.T) {
	type args struct {
		r udpRouterRepresentation
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_priority_0",
			args: args{
				r: udpRouterRepresentation{
					Name:     "test_priority_0",
					Provider: "test_priority_0",
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.r.priority(); got != tt.want {
				t.Errorf("priority() = %v, want %v", got, tt.want)
			}
		})
	}
}
