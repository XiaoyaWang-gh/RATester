package page_generate

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/codegen"
)

func TestGenerateMarshalJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &codegen.Inspector{}
	err := generateMarshalJSON(c)
	if err != nil {
		t.Fatal(err)
	}
}
