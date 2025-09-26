package gin

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin/binding"
)

func TestEnableJsonDecoderUseNumber_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	EnableJsonDecoderUseNumber()
	if !binding.EnableDecoderUseNumber {
		t.Error("Expected EnableDecoderUseNumber to be true, but it was false")
	}
}
