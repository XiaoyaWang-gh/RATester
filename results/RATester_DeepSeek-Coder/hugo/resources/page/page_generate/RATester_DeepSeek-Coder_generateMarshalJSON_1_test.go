package page_generate

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/codegen"
)

func TestGenerateMarshalJSON_1(t *testing.T) {
	type args struct {
		c *codegen.Inspector
	}
	tests := []struct {
		name    string
		args    args
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

			if err := generateMarshalJSON(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("generateMarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
