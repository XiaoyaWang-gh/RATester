package allconfig

import (
	"fmt"
	"testing"
)

func TestIsKindEnabled_1(t *testing.T) {
	type fields struct {
		C *ConfigCompiled
	}
	type args struct {
		kind string
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

			c := &Config{
				C: tt.fields.C,
			}
			if got := c.IsKindEnabled(tt.args.kind); got != tt.want {
				t.Errorf("IsKindEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
