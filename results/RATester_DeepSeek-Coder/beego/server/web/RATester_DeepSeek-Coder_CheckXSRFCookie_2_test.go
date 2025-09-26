package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/session"
)

func TestCheckXSRFCookie_2(t *testing.T) {
	type fields struct {
		Ctx        *context.Context
		Data       map[interface{}]interface{}
		EnableXSRF bool
		_xsrfToken string
		XSRFExpire int
		CruSession session.Store
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
				Ctx:        tt.fields.Ctx,
				Data:       tt.fields.Data,
				EnableXSRF: tt.fields.EnableXSRF,
				_xsrfToken: tt.fields._xsrfToken,
				XSRFExpire: tt.fields.XSRFExpire,
				CruSession: tt.fields.CruSession,
			}
			if got := c.CheckXSRFCookie(); got != tt.want {
				t.Errorf("CheckXSRFCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
