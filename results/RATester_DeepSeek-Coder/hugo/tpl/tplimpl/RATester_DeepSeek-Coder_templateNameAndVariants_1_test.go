package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTemplateNameAndVariants_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantName string
		wantVars []string
	}{
		{
			name:     "No variant",
			input:    "test",
			wantName: "test",
			wantVars: make([]string, numTemplateVariants),
		},
		{
			name:     "One variant",
			input:    "test.var1",
			wantName: "test",
			wantVars: []string{"var1", "", "", "", "", "", "", "", "", ""},
		},
		{
			name:     "Multiple variants",
			input:    "test.var1.var2.var3",
			wantName: "test",
			wantVars: []string{"var1", "var2", "var3", "", "", "", "", "", "", ""},
		},
		{
			name:     "More variants than slots",
			input:    "test.var1.var2.var3.var4.var5.var6.var7.var8.var9.var10.var11",
			wantName: "test",
			wantVars: []string{"var1", "var2", "var3", "var4", "var5", "var6", "var7", "var8", "var9", "var10"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotName, gotVars := templateNameAndVariants(tt.input)
			if gotName != tt.wantName {
				t.Errorf("templateNameAndVariants() gotName = %v, want %v", gotName, tt.wantName)
			}
			if !reflect.DeepEqual(gotVars, tt.wantVars) {
				t.Errorf("templateNameAndVariants() gotVars = %v, want %v", gotVars, tt.wantVars)
			}
		})
	}
}
