package allconfig

import (
	"fmt"
	"testing"
)

func TestDeleteMergeStrategies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A test context
	l := configLoader{}

	// when: The method under test is called
	l.deleteMergeStrategies()

	// then: ...
	// ...
}
