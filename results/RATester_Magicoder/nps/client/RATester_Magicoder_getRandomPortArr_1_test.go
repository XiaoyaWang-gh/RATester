package client

import (
	"fmt"
	"sort"
	"testing"
)

func TestgetRandomPortArr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	min := 1000
	max := 2000
	portArr := getRandomPortArr(min, max)

	if len(portArr) != max-min+1 {
		t.Errorf("Expected length of portArr is %d, but got %d", max-min+1, len(portArr))
	}

	for i := 0; i < len(portArr); i++ {
		if portArr[i] < min || portArr[i] > max {
			t.Errorf("Port number %d is out of range [%d, %d]", portArr[i], min, max)
		}
	}

	sort.Ints(portArr)
	for i := 0; i < len(portArr)-1; i++ {
		if portArr[i] == portArr[i+1] {
			t.Errorf("Duplicate port number %d found", portArr[i])
		}
	}
}
