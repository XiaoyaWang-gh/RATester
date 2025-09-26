package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateSectionHandlers_1(t *testing.T) {
	type args struct {
		l *pageLexer
	}
	tests := []struct {
		name string
		args args
		want *sectionHandlers
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

			if got := createSectionHandlers(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSectionHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
