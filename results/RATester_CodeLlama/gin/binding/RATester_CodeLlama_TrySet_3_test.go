package binding

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrySet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &multipartRequest{}
	r.MultipartForm = &multipart.Form{}
	r.MultipartForm.File = map[string][]*multipart.FileHeader{"key": {}}
	field := reflect.StructField{}
	opt := setOptions{}

	// when
	_, err := r.TrySet(reflect.Value{}, field, "key", opt)

	// then
	assert.NoError(t, err)
}
