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
		t.Errorf("EnableJsonDecoderUseNumber() failed, expected true but got false")
	}
}
