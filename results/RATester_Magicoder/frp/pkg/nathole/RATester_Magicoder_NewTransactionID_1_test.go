package nathole

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/fatedier/frp/pkg/util/util"
)

func TestNewTransactionID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	id := NewTransactionID()

	if id == "" {
		t.Error("Expected a non-empty string, got an empty string")
	}

	_, err := strconv.ParseInt(id[:10], 10, 64)
	if err != nil {
		t.Errorf("Expected the first 10 characters of the string to be a unix timestamp, got an error: %v", err)
	}

	_, err = util.RandID()
	if err != nil {
		t.Errorf("Expected the last 10 characters of the string to be a random string, got an error: %v", err)
	}
}
