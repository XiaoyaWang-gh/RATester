package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestStoreClientsToJsonFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	//given
	var jsonDb = &JsonDb{
		Clients: sync.Map{},
	}
	//when
	jsonDb.StoreClientsToJsonFile()
	//then
	//TODO
}
