package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInitExceptMethod_1(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test case 1",
			want: []string{"initExceptMethod", "Controller"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := initExceptMethod(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initExceptMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
