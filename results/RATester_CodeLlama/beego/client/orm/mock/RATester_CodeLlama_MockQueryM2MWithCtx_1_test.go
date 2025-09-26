package mock

import (
	"fmt"
	"testing"
)

func TestMockQueryM2MWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tableName := "post"
	name := "Tags"
	res := &DoNothingQueryM2Mer{}
	mock := MockQueryM2MWithCtx(tableName, name, res)
	if mock == nil {
		t.Errorf("MockQueryM2MWithCtx failed")
	}
}
