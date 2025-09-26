package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetBufPoolCopy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var bufPoolCopy []byte
	var poolSizeCopy int
	var expected []byte
	var actual []byte

	// Act
	bufPoolCopy = GetBufPoolCopy()
	poolSizeCopy = PoolSizeCopy
	expected = make([]byte, poolSizeCopy)
	actual = bufPoolCopy

	// Assert
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}
