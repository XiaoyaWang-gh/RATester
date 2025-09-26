package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestProtocol_1(t *testing.T) {
	type fields struct {
		Context *Context
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestProtocol",
			fields: fields{
				Context: &Context{
					Request: &http.Request{
						Proto: "HTTP/1.1",
					},
				},
			},
			want: "HTTP/1.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			input := &BeegoInput{
				Context: tt.fields.Context,
			}
			if got := input.Protocol(); got != tt.want {
				t.Errorf("BeegoInput.Protocol() = %v, want %v", got, tt.want)
			}
		})
	}
}
