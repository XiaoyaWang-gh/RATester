package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestsortClientByKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testMap := sync.Map{}
	testMap.Store("key1", &Client{Id: 1})
	testMap.Store("key2", &Client{Id: 2})
	testMap.Store("key3", &Client{Id: 3})

	res := sortClientByKey(testMap, "Id", "asc")

	if len(res) != 3 {
		t.Errorf("Expected length of result slice to be 3, but got %d", len(res))
	}

	if res[0] != 1 || res[1] != 2 || res[2] != 3 {
		t.Errorf("Expected sorted slice to be [1, 2, 3], but got %v", res)
	}
}
