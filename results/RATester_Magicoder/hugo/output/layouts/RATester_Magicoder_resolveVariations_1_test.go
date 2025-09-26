package layouts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestresolveVariations_1(t *testing.T) {
	type args struct {
		l *layoutBuilder
	}
	tests := []struct {
		name string
		args args
		want []string
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

			if got := tt.args.l.resolveVariations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveVariations() = %v, want %v", got, tt.want)
			}
		})
	}
}
