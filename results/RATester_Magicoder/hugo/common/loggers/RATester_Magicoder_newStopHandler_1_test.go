package loggers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestnewStopHandler_1(t *testing.T) {
	type args struct {
		h []logg.Handler
	}
	tests := []struct {
		name string
		args args
		want *stopHandler
	}{
		{
			name: "Test Case 1",
			args: args{
				h: []logg.Handler{
					// Add your handlers here
				},
			},
			want: &stopHandler{
				handlers: []logg.Handler{
					// Add your handlers here
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

			if got := newStopHandler(tt.args.h...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStopHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
