package web

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/session"
)

func TestStartSession_1(t *testing.T) {
	type fields struct {
		Ctx        *context.Context
		Data       map[interface{}]interface{}
		CruSession session.Store
	}
	tests := []struct {
		name   string
		fields fields
		want   session.Store
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
				CruSession: tt.fields.CruSession,
			}
			if got := c.StartSession(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.StartSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
