package nomad

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTagsToLabels_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tags := []string{"traefik.frontend.rule=Host:test.traefik.io", "traefik.port=80", "other=test"}
	prefix := "traefik.frontend"
	expected := map[string]string{"traefik.rule": "Host:test.traefik.io"}
	actual := tagsToLabels(tags, prefix)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
