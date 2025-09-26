package buffers

import (
	"fmt"
	"testing"

	"github.com/valyala/bytebufferpool"
)

func TestPut_8(t *testing.T) {
	type args struct {
		bf Buffer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestPut",
			args: args{
				bf: &bytebufferpool.ByteBuffer{},
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

			Put(tt.args.bf)
		})
	}
}
