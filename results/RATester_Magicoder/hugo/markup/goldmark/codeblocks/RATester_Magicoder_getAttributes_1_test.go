package codeblocks

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestgetAttributes_1(t *testing.T) {
	type args struct {
		node    *ast.FencedCodeBlock
		infostr []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []ast.Attribute
		want1   string
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

			got, got1, err := getAttributes(tt.args.node, tt.args.infostr)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAttributes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAttributes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getAttributes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
