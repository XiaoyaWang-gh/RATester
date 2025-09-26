package passthrough

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/internal/attributes"
)

func TestType_1(t *testing.T) {
	type fields struct {
		typ   string
		inner string
		attrs *attributes.AttributesHolder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestType",
			fields: fields{
				typ:   "inner",
				inner: "inner",
				attrs: &attributes.AttributesHolder{},
			},
			want: "inner",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &passthroughContext{
				typ:              tt.fields.typ,
				inner:            tt.fields.inner,
				AttributesHolder: tt.fields.attrs,
			}
			if got := p.Type(); got != tt.want {
				t.Errorf("passthroughContext.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
