package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTagsToNeutralLabels_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tags := []string{"traefik.frontend.rule=Host:test.traefik.io", "traefik.port=80"}
	prefix := "traefik"
	expected := map[string]string{"traefik.frontend.rule": "Host:test.traefik.io", "traefik.port": "80"}
	actual := tagsToNeutralLabels(tags, prefix)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
