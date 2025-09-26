package framework

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExplain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &RequestExpect{}
	explain := []interface{}{"This is a test explain"}
	e.Explain(explain...)
	if !reflect.DeepEqual(e.explain, explain) {
		t.Errorf("Expected explain to be %v, but got %v", explain, e.explain)
	}
}
