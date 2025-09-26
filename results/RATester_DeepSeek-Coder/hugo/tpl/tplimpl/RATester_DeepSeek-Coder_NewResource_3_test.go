package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNewResource_3(t *testing.T) {
	type args struct {
		dst *deps.Deps
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestNewResource",
			args: args{
				dst: &deps.Deps{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tp := &TemplateProvider{}
			if err := tp.NewResource(tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("TemplateProvider.NewResource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
