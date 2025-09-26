package web

import (
	"fmt"
	"strings"
	"testing"
)

func TestQpsIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &adminController{}
	a.QpsIndex()

	// Test case 1: Check if the data map is correctly initialized
	if a.Data == nil {
		t.Error("Data map is not initialized")
	}

	// Test case 2: Check if the Content key in the data map is correctly set
	content, ok := a.Data["Content"].(M)
	if !ok {
		t.Error("Content key in the data map is not of type M")
	}

	// Test case 3: Check if the Data key in the content map is correctly set
	resultLists, ok := content["Data"].([][]string)
	if !ok {
		t.Error("Data key in the content map is not of type [][]string")
	}

	// Test case 4: Check if the first element of each sub-slice is correctly escaped
	for i := range resultLists {
		if len(resultLists[i]) > 0 {
			if !strings.HasPrefix(resultLists[i][0], "&lt;") || !strings.HasSuffix(resultLists[i][0], "&gt;") {
				t.Errorf("First element of sub-slice %d is not correctly escaped", i)
			}
		}
	}
}
