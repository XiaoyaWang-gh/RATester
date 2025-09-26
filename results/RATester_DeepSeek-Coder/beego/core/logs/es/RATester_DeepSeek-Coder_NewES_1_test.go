package es

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func TestNewES_1(t *testing.T) {
	tests := []struct {
		name string
		want logs.Logger
	}{
		{
			name: "TestNewES",
			want: &esLogger{
				Level:       logs.LevelDebug,
				indexNaming: indexNaming,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewES(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewES() = %v, want %v", got, tt.want)
			}
		})
	}
}
