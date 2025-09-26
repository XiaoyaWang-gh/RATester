package i18n

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNewResource_2(t *testing.T) {
	type args struct {
		dst *deps.Deps
	}
	tests := []struct {
		name    string
		tp      *TranslationProvider
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

			if err := tt.tp.NewResource(tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("TranslationProvider.NewResource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
