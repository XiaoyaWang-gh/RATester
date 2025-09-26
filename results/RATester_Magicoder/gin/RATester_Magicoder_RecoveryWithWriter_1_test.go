package gin

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestRecoveryWithWriter_1(t *testing.T) {
	type args struct {
		out      io.Writer
		recovery []RecoveryFunc
	}
	tests := []struct {
		name string
		args args
		want HandlerFunc
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

			if got := RecoveryWithWriter(tt.args.out, tt.args.recovery...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecoveryWithWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
