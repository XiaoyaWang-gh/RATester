package session

import (
	"fmt"
	"testing"
)

func TestRegisterType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var i any
	// act
	s := &Store{}
	s.RegisterType(i)
	// assert
	// TODO: add assertions
}
