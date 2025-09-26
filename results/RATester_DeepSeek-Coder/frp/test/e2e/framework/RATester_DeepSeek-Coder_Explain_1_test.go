package framework

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestExplain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &RequestExpect{
		req: &request.Request{},
	}

	explain := []interface{}{"test", 123, true}
	e.Explain(explain...)

	if !reflect.DeepEqual(e.explain, explain) {
		t.Errorf("Expected explain to be %v, but got %v", explain, e.explain)
	}
}
