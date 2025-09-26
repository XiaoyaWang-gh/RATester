package echo

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	w := httptest.NewRecorder()
	e := &Echo{}

	// act
	r := NewResponse(w, e)

	// assert
	assert.NotNil(t, r)
	assert.Equal(t, w, r.Writer)
	assert.Equal(t, e, r.echo)
}
