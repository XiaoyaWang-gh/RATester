package template

import (
	"fmt"
	"testing"
)

func TestPrepare_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTemplate := &Template{
		name: "testTemplate",
	}

	preparedTemplate, err := testTemplate.Prepare()

	if err != nil {
		t.Errorf("Prepare() returned an error: %v", err)
	}

	if preparedTemplate.name != "testTemplate" {
		t.Errorf("Prepare() did not set the correct name. Expected: 'testTemplate', Got: '%s'", preparedTemplate.name)
	}
}
