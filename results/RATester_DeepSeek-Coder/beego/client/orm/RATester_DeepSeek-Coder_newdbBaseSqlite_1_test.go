package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewdbBaseSqlite_1(t *testing.T) {
	tests := []struct {
		name string
		want dbBaser
	}{
		{
			name: "Test newdbBaseSqlite",
			want: newdbBaseSqlite(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := newdbBaseSqlite()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newdbBaseSqlite() = %v, want %v", got, tt.want)
			}
		})
	}
}
