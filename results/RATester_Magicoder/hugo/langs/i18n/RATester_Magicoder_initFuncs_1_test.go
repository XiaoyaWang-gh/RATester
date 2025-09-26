package i18n

import (
	"fmt"
	"testing"

	"github.com/gohugoio/go-i18n/v2/i18n"
)

func TestinitFuncs_1(t *testing.T) {
	type args struct {
		bndl *i18n.Bundle
	}
	tests := []struct {
		name string
		args args
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

			tr := Translator{}
			tr.initFuncs(tt.args.bndl)
		})
	}
}
