package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAsync_1(t *testing.T) {
	tests := []struct {
		name   string
		msgLen []int64
		want   *BeeLogger
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

			if got := Async(tt.msgLen...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Async() = %v, want %v", got, tt.want)
			}
		})
	}
}
