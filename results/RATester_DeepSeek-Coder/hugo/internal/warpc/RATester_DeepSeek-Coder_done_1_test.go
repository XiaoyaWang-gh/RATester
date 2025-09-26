package warpc

import (
	"fmt"
	"testing"
)

func TestDone_1(t *testing.T) {
	type args struct {
		call *call[int, string]
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				call: &call[int, string]{
					donec: make(chan *call[int, string]),
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

			tt.args.call.done()
			select {
			case <-tt.args.call.donec:
				// Test passed
			default:
				t.Errorf("done() did not close donec")
			}
		})
	}
}
