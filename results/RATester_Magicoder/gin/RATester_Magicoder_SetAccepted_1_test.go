package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetAccepted_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{}
	formats := []string{"format1", "format2"}
	ctx.SetAccepted(formats...)

	if !reflect.DeepEqual(ctx.Accepted, formats) {
		t.Errorf("Expected Accepted to be %v, but got %v", formats, ctx.Accepted)
	}
}
