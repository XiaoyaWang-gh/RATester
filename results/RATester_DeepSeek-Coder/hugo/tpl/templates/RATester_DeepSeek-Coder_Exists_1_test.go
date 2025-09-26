package templates

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestExists_1(t *testing.T) {
	type fields struct {
		deps *deps.Deps
	}
	type args struct {
		name string
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

			ns := &Namespace{
				deps: tt.fields.deps,
			}
			if got := ns.Exists(tt.args.name); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
