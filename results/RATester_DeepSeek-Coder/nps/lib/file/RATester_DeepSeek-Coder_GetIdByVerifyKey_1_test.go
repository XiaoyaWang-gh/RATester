package file

import (
	"fmt"
	"testing"
)

func TestGetIdByVerifyKey_1(t *testing.T) {
	db := &DbUtils{
		JsonDb: &JsonDb{},
	}

	t.Run("TestGetIdByVerifyKey", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := db.GetIdByVerifyKey("vKey", "addr")
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
}
