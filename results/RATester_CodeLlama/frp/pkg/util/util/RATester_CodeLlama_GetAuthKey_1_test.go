package util

import (
	"fmt"
	"testing"
)

func TestGetAuthKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	token := "1234567890"
	timestamp := int64(1589888000)
	key := GetAuthKey(token, timestamp)
	if key != "1234567890" {
		t.Errorf("GetAuthKey() = %v, want %v", key, "1234567890")
	}
}
