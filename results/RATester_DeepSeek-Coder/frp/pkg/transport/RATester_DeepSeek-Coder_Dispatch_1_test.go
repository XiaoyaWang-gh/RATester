package transport

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestDispatch_1(t *testing.T) {
	type args struct {
		m       msg.Message
		laneKey string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			impl := &transporterImpl{}
			if got := impl.Dispatch(tt.args.m, tt.args.laneKey); got != tt.want {
				t.Errorf("transporterImpl.Dispatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
