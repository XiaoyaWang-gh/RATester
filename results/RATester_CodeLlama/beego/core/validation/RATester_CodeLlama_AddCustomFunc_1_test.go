package validation

import (
	"fmt"
	"testing"

	"github.com/beego/beego/validation"
	"github.com/zeebo/assert"
)

func TestAddCustomFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	name := "test"
	f := func(v *validation.Validation, obj interface{}, key string) {}

	// when
	err := validation.AddCustomFunc(name, f)

	// then
	assert.NoError(t, err)
}
