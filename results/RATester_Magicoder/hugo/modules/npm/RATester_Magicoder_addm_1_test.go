package npm

import (
	"fmt"
	"reflect"
	"testing"
)

func Testaddm_1(t *testing.T) {
	type args struct {
		source string
		m      map[string]any
	}
	tests := []struct {
		name string
		args args
		want map[string]any
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

			b := &packageBuilder{}
			b.addm(tt.args.source, tt.args.m)
			if !reflect.DeepEqual(b.devDependencies, tt.want) {
				t.Errorf("addm() = %v, want %v", b.devDependencies, tt.want)
			}
		})
	}
}
