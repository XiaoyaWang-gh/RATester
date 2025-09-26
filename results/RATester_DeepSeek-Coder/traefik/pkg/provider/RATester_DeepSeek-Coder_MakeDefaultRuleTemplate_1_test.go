package provider

import (
	"fmt"
	"reflect"
	"testing"
	"text/template"
)

func TestMakeDefaultRuleTemplate_1(t *testing.T) {
	type args struct {
		defaultRule string
		funcMap     template.FuncMap
	}
	tests := []struct {
		name    string
		args    args
		want    *template.Template
		wantErr bool
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

			got, err := MakeDefaultRuleTemplate(tt.args.defaultRule, tt.args.funcMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeDefaultRuleTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeDefaultRuleTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
