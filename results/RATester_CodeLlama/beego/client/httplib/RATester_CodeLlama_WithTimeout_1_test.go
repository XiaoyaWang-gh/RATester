package httplib

import (
	"fmt"
	"testing"
	"time"
)

func TestWithTimeout_1(t *testing.T) {
	type args struct {
		connectTimeout   time.Duration
		readWriteTimeout time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				connectTimeout:   10 * time.Second,
				readWriteTimeout: 10 * time.Second,
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

			WithTimeout(tt.args.connectTimeout, tt.args.readWriteTimeout)
		})
	}
}
