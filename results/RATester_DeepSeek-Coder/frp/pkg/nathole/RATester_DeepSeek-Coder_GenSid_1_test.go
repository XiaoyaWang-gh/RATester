package nathole

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/google/uuid"
)

func TestGenSid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{}
	sid := ctrl.GenSid()

	if sid == "" {
		t.Errorf("Expected a non-empty string, got an empty string")
	}

	_, err := strconv.ParseInt(sid[:10], 10, 64)
	if err != nil {
		t.Errorf("Expected the first 10 characters to be a unix timestamp, got %v", err)
	}

	_, err = uuid.Parse(sid[10:])
	if err != nil {
		t.Errorf("Expected the last characters to be a valid UUID, got %v", err)
	}
}
