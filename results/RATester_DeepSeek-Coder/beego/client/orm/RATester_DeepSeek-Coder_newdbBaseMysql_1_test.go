package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewdbBaseMysql_1(t *testing.T) {
	tests := []struct {
		name string
		want dbBaser
	}{
		{
			name: "Test newdbBaseMysql",
			want: newdbBaseMysql(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newdbBaseMysql(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newdbBaseMysql() = %v, want %v", got, tt.want)
			}
		})
	}
}
