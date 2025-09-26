package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetOneOPage_1(t *testing.T) {
	type args struct {
		t OrderedTaxonomy
	}
	tests := []struct {
		name string
		args args
		want Page
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

			if got := tt.args.t.getOneOPage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOneOPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
