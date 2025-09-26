package echo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/zeebo/assert"
)

func TestBindBody_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	b := &DefaultBinder{}
	c := &context{
		request: &http.Request{
			Header: http.Header{
				HeaderContentType: []string{MIMEApplicationJSON},
			},
			Body: ioutil.NopCloser(strings.NewReader(`{"name":"john","age":30}`)),
		},
	}
	i := &struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	// When
	err := b.BindBody(c, i)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "john", i.Name)
	assert.Equal(t, 30, i.Age)
}
