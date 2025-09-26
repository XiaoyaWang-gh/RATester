package render

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	w := httptest.NewRecorder()
	r := IndentedJSON{Data: "test"}

	// when
	err := r.Render(w)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "{\n    \"Data\": \"test\"\n}", w.Body.String())
}
