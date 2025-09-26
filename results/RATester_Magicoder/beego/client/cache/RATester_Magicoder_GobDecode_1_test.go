package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGobDecode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := &FileCacheItem{
		Data:       "test data",
		Lastaccess: time.Now(),
		Expired:    time.Now().Add(time.Hour),
	}

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(item)
	if err != nil {
		t.Fatalf("Failed to encode item: %v", err)
	}

	decodedItem := &FileCacheItem{}
	err = GobDecode(buf.Bytes(), decodedItem)
	if err != nil {
		t.Fatalf("Failed to decode item: %v", err)
	}

	if !reflect.DeepEqual(item, decodedItem) {
		t.Errorf("Decoded item does not match original item")
	}
}
