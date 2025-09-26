package attributes

import (
	"fmt"
	"strings"
	"testing"
)

func Testinit_31(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	chromaHighlightProcessingAttributes := map[string]string{
		"key1": "value1",
		"key2": "value2",
		// add more test data as needed
	}

	for k, v := range chromaHighlightProcessingAttributes {
		chromaHighlightProcessingAttributes[strings.ToLower(k)] = v
	}

	// check if the map has been modified
	for k, v := range chromaHighlightProcessingAttributes {
		if strings.ToLower(k) != k {
			t.Errorf("Key %s was not converted to lowercase", k)
		}
		if v != chromaHighlightProcessingAttributes[strings.ToLower(k)] {
			t.Errorf("Value for key %s was not preserved", k)
		}
	}
}
