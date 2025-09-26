package binding

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBind_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	req := httptest.NewRequest("POST", "/", strings.NewReader("name=john&age=20"))
	obj := &struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}{}

	// when
	err := (&plainBinding{}).Bind(req, obj)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "john", obj.Name)
	assert.Equal(t, 20, obj.Age)
}
