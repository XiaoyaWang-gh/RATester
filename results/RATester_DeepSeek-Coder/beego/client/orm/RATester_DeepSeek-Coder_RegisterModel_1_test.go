package orm

import (
	"fmt"
	"testing"
)

func TestRegisterModel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		ID   int
		Name string
	}

	RegisterModel(TestStruct{})

	// Add more assertions here to test the functionality of the RegisterModel function
}
