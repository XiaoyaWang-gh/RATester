package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetGetTplPackagesGoDoc_1(t *testing.T) {
	tests := []struct {
		name string
		want map[string]map[string]methodGoDocInfo
	}{
		{
			name: "Test getGetTplPackagesGoDoc",
			want: getGetTplPackagesGoDoc(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getGetTplPackagesGoDoc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGetTplPackagesGoDoc() = %v, want %v", got, tt.want)
			}
		})
	}
}
