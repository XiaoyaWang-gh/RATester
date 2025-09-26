package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/config"
)

func TestDefaultFloat_8(t *testing.T) {
	type fields struct {
		BaseConfiger config.BaseConfiger
		innerConfig  config.Configer
	}
	type args struct {
		key        string
		defaultVal float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
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

			b := &beegoAppConfig{
				BaseConfiger: tt.fields.BaseConfiger,
				innerConfig:  tt.fields.innerConfig,
			}
			if got := b.DefaultFloat(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("beegoAppConfig.DefaultFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
