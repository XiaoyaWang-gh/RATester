package gin

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin/binding"
)

func TestMustBindWith_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	var c Context
	var obj any
	var b binding.Binding
	var err error

	err = c.MustBindWith(obj, b)
	if err != nil {
		t.Errorf("MustBindWith() error = %v, want nil", err)
		return
	}
}
