package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestIsAjax_2(t *testing.T) {
	type fields struct {
		Ctx  *context.Context
		Data map[interface{}]interface{}
	}
	tests := []struct {
		name   string
		fields fields
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

			c := &Controller{
				Ctx:  tt.fields.Ctx,
				Data: tt.fields.Data,
			}
			if got := c.IsAjax(); got != tt.want {
				t.Errorf("IsAjax() = %v, want %v", got, tt.want)
			}
		})
	}
}
