package gin

import (
	"fmt"
	"testing"
)

func TestBindXML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
