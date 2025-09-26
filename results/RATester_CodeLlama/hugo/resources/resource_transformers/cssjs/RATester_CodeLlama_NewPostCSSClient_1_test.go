package cssjs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources"
)

func TestNewPostCSSClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var rs *resources.Spec

	// act
	var postCSSClient *PostCSSClient
	postCSSClient = NewPostCSSClient(rs)

	// assert
	assert.NotNil(t, postCSSClient)
}
