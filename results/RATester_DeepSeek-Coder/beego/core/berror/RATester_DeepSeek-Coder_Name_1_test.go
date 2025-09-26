package berror

import (
	"fmt"
	"testing"
)

func TestName_1(t *testing.T) {
	type fields struct {
		code   uint32
		module string
		desc   string
		name   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestName",
			fields: fields{
				code:   1,
				module: "module",
				desc:   "desc",
				name:   "name",
			},
			want: "name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &codeDefinition{
				code:   tt.fields.code,
				module: tt.fields.module,
				desc:   tt.fields.desc,
				name:   tt.fields.name,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("codeDefinition.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
