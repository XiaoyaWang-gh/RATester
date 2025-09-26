package plugins

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPpSymbols_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	got := ppSymbols()
	want := map[string]map[string]reflect.Value{
		"github.com/traefik/traefik/v3/pkg/plugins/plugins": {
			"PP":  reflect.ValueOf((*PP)(nil)),
			"_PP": reflect.ValueOf((*_PP)(nil)),
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ppSymbols() = %v, want %v", got, want)
	}
}
