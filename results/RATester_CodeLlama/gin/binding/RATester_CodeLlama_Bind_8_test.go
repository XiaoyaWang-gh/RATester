package binding

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBind_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	req := httptest.NewRequest("POST", "/", strings.NewReader(`<xml><name>John</name><age>25</age></xml>`))
	var obj struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}

	// when
	err := xmlBinding{}.Bind(req, &obj)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "John", obj.Name)
	assert.Equal(t, 25, obj.Age)
}
