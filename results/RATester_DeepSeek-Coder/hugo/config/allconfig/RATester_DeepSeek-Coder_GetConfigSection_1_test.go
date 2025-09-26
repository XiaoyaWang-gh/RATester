package allconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetConfigSection_1(t *testing.T) {
	type fields struct {
		config *Config
		m      *Configs
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
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

			c := ConfigLanguage{
				config: tt.fields.config,
				m:      tt.fields.m,
			}
			if got := c.GetConfigSection(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfigSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
