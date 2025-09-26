package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeByLanguage_2(t *testing.T) {
	type args struct {
		p1 Pages
		p2 Pages
	}
	tests := []struct {
		name string
		args args
		want Pages
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

			if got := tt.args.p1.MergeByLanguage(tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeByLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}
