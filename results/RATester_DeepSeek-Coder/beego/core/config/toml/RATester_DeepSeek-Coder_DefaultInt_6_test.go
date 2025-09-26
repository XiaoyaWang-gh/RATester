package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestDefaultInt_6(t *testing.T) {
	type fields struct {
		t *toml.Tree
	}
	type args struct {
		key        string
		defaultVal int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
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

			c := &configContainer{
				t: tt.fields.t,
			}
			if got := c.DefaultInt(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
