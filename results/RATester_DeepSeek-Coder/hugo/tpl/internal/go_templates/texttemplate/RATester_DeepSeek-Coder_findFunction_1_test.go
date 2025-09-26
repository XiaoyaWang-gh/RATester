package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindFunction_1(t *testing.T) {
	type args struct {
		name string
		tmpl *Template
	}
	tests := []struct {
		name          string
		args          args
		wantV         reflect.Value
		wantIsBuiltin bool
		wantOk        bool
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

			gotV, gotIsBuiltin, gotOk := findFunction(tt.args.name, tt.args.tmpl)
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("findFunction() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotIsBuiltin != tt.wantIsBuiltin {
				t.Errorf("findFunction() gotIsBuiltin = %v, want %v", gotIsBuiltin, tt.wantIsBuiltin)
			}
			if gotOk != tt.wantOk {
				t.Errorf("findFunction() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
