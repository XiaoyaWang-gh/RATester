package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestlookupAndEscapeTemplate_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		t        *Template
		args     args
		wantTmpl *Template
		wantErr  bool
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

			gotTmpl, err := tt.t.lookupAndEscapeTemplate(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.lookupAndEscapeTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTmpl, tt.wantTmpl) {
				t.Errorf("Template.lookupAndEscapeTemplate() = %v, want %v", gotTmpl, tt.wantTmpl)
			}
		})
	}
}
