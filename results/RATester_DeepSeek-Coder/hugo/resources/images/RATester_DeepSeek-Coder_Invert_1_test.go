package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/disintegration/gift"
)

func TestInvert_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}
	invertFilter := filters.Invert()

	if reflect.TypeOf(invertFilter) != reflect.TypeOf(gift.Invert()) {
		t.Errorf("Expected type %v, got %v", reflect.TypeOf(gift.Invert()), reflect.TypeOf(invertFilter))
	}
}
