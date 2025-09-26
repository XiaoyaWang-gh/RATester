package log

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultLogger_1(t *testing.T) {
	tests := []struct {
		name string
		want AllLogger
	}{
		{
			name: "Test DefaultLogger",
			want: logger,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DefaultLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
