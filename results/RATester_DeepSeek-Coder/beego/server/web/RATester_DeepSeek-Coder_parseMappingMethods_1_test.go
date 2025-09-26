package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseMappingMethods_1(t *testing.T) {
	type args struct {
		c              ControllerInterface
		mappingMethods []string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := parseMappingMethods(tt.args.c, tt.args.mappingMethods); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMappingMethods() = %v, want %v", got, tt.want)
			}
		})
	}
}
