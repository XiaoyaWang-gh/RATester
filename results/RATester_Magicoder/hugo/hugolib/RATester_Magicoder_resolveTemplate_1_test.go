package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func TestresolveTemplate_1(t *testing.T) {
	type args struct {
		layouts []string
	}
	tests := []struct {
		name    string
		p       *pageState
		args    args
		want    tpl.Template
		want1   bool
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

			got, got1, err := tt.p.resolveTemplate(tt.args.layouts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("resolveTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveTemplate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("resolveTemplate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
