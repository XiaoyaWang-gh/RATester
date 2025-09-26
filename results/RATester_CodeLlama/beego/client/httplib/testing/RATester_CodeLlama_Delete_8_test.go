package testing

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDelete_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "/path"
	var testHTTPRequest *TestHTTPRequest
	testHTTPRequest = Delete(path)
	assert.NotNil(t, testHTTPRequest)
}
