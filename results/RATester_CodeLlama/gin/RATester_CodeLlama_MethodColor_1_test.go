package gin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/assert"
)

func TestMethodColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &LogFormatterParams{
		Method: http.MethodGet,
	}
	assert.Equal(t, blue, p.MethodColor())
}
