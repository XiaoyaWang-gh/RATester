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

			tp := &TranslationProvider{}
			if err := tp.NewResource(tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("TranslationProvider.NewResource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
