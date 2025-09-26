package server

import (
	"fmt"
	"testing"
)

func TestGetClientList_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	start := 0
	length := 10
	search := ""
	sort := "id"
	order := "asc"
	clientId := 1

	list, cnt := GetClientList(start, length, search, sort, order, clientId)

	if len(list) != cnt {
		t.Errorf("Expected length of list to be equal to count, but got %d and %d", len(list), cnt)
	}

	for _, client := range list {
		if client.Id != clientId {
			t.Errorf("Expected client id to be %d, but got %d", clientId, client.Id)
		}
	}
}
