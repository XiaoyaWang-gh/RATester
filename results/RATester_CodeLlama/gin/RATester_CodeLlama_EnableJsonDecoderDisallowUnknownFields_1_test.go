package gin

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin/binding"
)

func TestEnableJsonDecoderDisallowUnknownFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	EnableJsonDecoderDisallowUnknownFields()
	if binding.EnableDecoderDisallowUnknownFields != true {
		t.Errorf("EnableJsonDecoderDisallowUnknownFields() failed")
	}
}
