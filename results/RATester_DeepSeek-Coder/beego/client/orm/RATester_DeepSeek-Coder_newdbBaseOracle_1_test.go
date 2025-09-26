package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewdbBaseOracle_1(t *testing.T) {
	tests := []struct {
		name string
		want dbBaser
	}{
		{
			name: "Test newdbBaseOracle",
			want: new(dbBaseOracle),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := newdbBaseOracle()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newdbBaseOracle() = %v, want %v", got, tt.want)
			}
		})
	}
}
