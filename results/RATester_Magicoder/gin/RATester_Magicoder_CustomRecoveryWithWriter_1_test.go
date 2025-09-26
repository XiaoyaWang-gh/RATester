package gin

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestCustomRecoveryWithWriter_1(t *testing.T) {
	tests := []struct {
		name   string
		out    io.Writer
		handle RecoveryFunc
		want   HandlerFunc
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

			if got := CustomRecoveryWithWriter(tt.out, tt.handle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomRecoveryWithWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
