package hugolib

import (
	"fmt"
	"testing"
)

func TestData_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageData{}
	p.Data()
}
