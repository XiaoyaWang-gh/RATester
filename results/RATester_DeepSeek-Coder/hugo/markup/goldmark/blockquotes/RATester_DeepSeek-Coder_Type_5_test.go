package blockquotes

import (
	"fmt"
	"testing"
)

func TestType_5(t *testing.T) {
	type fields struct {
		typ string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestType",
			fields: fields{
				typ: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &blockquoteContext{
				typ: tt.fields.typ,
			}
			if got := c.Type(); got != tt.want {
				t.Errorf("blockquoteContext.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
