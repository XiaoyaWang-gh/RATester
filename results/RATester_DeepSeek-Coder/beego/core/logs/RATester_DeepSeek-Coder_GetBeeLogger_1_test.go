package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetBeeLogger_1(t *testing.T) {
	tests := []struct {
		name string
		want *BeeLogger
	}{
		{
			name: "TestGetBeeLogger",
			want: GetBeeLogger(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetBeeLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBeeLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
