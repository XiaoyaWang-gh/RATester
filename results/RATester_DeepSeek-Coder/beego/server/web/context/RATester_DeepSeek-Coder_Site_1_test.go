package context

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/beego/beego/v2/server/web/session"
)

func TestSite_1(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
		want   string
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
			if got := input.Site(); got != tt.want {
				t.Errorf("BeegoInput.Site() = %v, want %v", got, tt.want)
			}
		})
	}
}
