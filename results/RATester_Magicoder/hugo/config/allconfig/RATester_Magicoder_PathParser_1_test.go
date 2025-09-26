package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
)

func TestPathParser_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	configLanguage := ConfigLanguage{
		m: &Configs{
			ContentPathParser: &paths.PathParser{},
		},
	}

	expected := &paths.PathParser{}
	actual := configLanguage.PathParser()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
