package render

import (
	"fmt"
	"html/template"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteContentType_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	w := httptest.NewRecorder()
	r := &HTML{
		Template: template.Must(template.New("").Parse("")),
		Name:     "test",
		Data:     nil,
	}

	// when
	r.WriteContentType(w)

	// then
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}
