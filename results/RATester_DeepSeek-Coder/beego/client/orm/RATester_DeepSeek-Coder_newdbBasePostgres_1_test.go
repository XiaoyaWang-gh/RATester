package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewdbBasePostgres_1(t *testing.T) {
	tests := []struct {
		name string
		want dbBaser
	}{
		{
			name: "Test newdbBasePostgres",
			want: new(dbBasePostgres),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := newdbBasePostgres()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newdbBasePostgres() = %v, want %v", got, tt.want)
			}
		})
	}
}
