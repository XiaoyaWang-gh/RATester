package client

import (
	"fmt"
	"testing"
)

func TestSetServiceOptionsDefault_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	options := &ServiceOptions{}
	setServiceOptionsDefault(options)
	if options.Common == nil {
		t.Error("options.Common should not be nil")
	}
	if options.ConnectorCreator == nil {
		t.Error("options.ConnectorCreator should not be nil")
	}
}
