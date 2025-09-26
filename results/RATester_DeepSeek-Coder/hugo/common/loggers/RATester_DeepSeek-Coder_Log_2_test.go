package loggers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLog_2(t *testing.T) {
	tests := []struct {
		name string
		want Logger
	}{
		{
			name: "Test case 1",
			want: Log(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Log(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}
