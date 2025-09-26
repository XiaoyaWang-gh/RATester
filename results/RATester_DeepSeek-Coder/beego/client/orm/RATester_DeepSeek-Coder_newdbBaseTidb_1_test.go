package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewdbBaseTidb_1(t *testing.T) {
	tests := []struct {
		name string
		want dbBaser
	}{
		{
			name: "Test newdbBaseTidb",
			want: new(dbBaseTidb),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := newdbBaseTidb()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newdbBaseTidb() = %v, want %v", got, tt.want)
			}
		})
	}
}
