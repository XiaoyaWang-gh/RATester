package binding

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBind_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	req := httptest.NewRequest("POST", "/", strings.NewReader("key = \"value\""))
	req.Header.Set("Content-Type", "application/toml")
	var obj map[string]string

	// when
	err := tomlBinding{}.Bind(req, &obj)

	// then
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"key": "value"}, obj)
}
