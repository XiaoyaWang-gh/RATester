package allconfig

import (
	"fmt"
	"testing"
)

func TestIgnoreFile_1(t *testing.T) {
	type fields struct {
		config *Config
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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
			}
			if got := c.IgnoreFile(tt.args.s); got != tt.want {
				t.Errorf("IgnoreFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
