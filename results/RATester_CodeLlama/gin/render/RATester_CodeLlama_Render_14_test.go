package render

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	pureJSON := PureJSON{
		Data: "test",
	}
	w := httptest.NewRecorder()
	// when
	err := pureJSON.Render(w)
	// then
	assert.NoError(t, err)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "test", w.Body.String())
}
