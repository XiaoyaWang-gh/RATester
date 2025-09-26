package logs

import (
	"fmt"
	"testing"
)

func TestGetLogFuncCallDepth_1(t *testing.T) {
	tests := []struct {
		name string
		bl   *BeeLogger
		want int
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

			if got := tt.bl.GetLogFuncCallDepth(); got != tt.want {
				t.Errorf("BeeLogger.GetLogFuncCallDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
