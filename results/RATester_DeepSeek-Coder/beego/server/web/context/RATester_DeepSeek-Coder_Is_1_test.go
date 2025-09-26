package context

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/beego/beego/v2/server/web/session"
)

func TestIs_1(t *testing.T) {
	type fields struct {
		Context       *Context
		CruSession    session.Store
		pnames        []string
		pvalues       []string
		data          map[interface{}]interface{}
		dataLock      sync.RWMutex
		RequestBody   []byte
		RunMethod     string
		RunController reflect.Type
	}
	type args struct {
		method string
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

			input := &BeegoInput{
				Context:       tt.fields.Context,
				CruSession:    tt.fields.CruSession,
				pnames:        tt.fields.pnames,
				pvalues:       tt.fields.pvalues,
				data:          tt.fields.data,
				dataLock:      tt.fields.dataLock,
				RequestBody:   tt.fields.RequestBody,
				RunMethod:     tt.fields.RunMethod,
				RunController: tt.fields.RunController,
			}
			if got := input.Is(tt.args.method); got != tt.want {
				t.Errorf("BeegoInput.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
