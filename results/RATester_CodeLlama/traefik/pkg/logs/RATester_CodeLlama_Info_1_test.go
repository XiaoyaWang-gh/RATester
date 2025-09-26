package logs

import (
	"fmt"
	"testing"
)

func TestInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var l OxyWrapper
	var s string
	var i []interface{}

	// act
	l.Info(s, i...)

	// assert
	// TODO
}
