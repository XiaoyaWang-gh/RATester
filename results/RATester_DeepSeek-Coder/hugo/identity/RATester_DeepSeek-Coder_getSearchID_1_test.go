package identity

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/compare"
)

func TestGetSearchID_1(t *testing.T) {
	type args struct {
		id       Identity
		isDp     bool
		isPeq    bool
		hasEqer  bool
		maxDepth int
		seen     map[Manager]bool
		dp       IsProbablyDependentProvider
		peq      compare.ProbablyEqer
		eqer     compare.Eqer
	}
	tests := []struct {
		name string
		args args
		want *searchID
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

			if got := getSearchID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSearchID() = %v, want %v", got, tt.want)
			}
		})
	}
}
