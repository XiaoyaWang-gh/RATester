package common

import (
	"fmt"
	"testing"
)

func TestPutBufPoolUdp_1(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				buf: make([]byte, PoolSizeUdp),
			},
		},
		{
			name: "Test case 2",
			args: args{
				buf: make([]byte, PoolSizeUdp+1),
			},
		},
		{
			name: "Test case 3",
			args: args{
				buf: make([]byte, PoolSizeUdp-1),
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

			PutBufPoolUdp(tt.args.buf)
		})
	}
}
