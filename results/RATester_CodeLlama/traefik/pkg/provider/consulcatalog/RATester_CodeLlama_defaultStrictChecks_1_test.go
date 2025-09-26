package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestDefaultStrictChecks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := []string{api.HealthPassing, api.HealthWarning}
	actual := defaultStrictChecks()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
