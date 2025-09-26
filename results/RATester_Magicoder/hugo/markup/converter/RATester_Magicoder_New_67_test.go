package converter

import (
	"fmt"
	"testing"
)

func TestNew_67(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := DocumentContext{
		Document:       nil,
		DocumentLookup: nil,
		DocumentID:     "testID",
		DocumentName:   "testName",
		Filename:       "testFilename",
	}

	n := newConverter{
		name: "testName",
		create: func(ctx DocumentContext) (Converter, error) {
			return nil, nil
		},
	}

	_, err := n.New(ctx)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
