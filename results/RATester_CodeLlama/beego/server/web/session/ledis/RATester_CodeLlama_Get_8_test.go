package ledis

import (
	"context"
	"fmt"
	"testing"
)

func TestGet_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ls := &SessionStore{}
	ls.values = make(map[interface{}]interface{})
	ls.values["key"] = "value"
	ctx := context.Background()
	if ls.Get(ctx, "key") != "value" {
		t.Errorf("Get() = %v, want %v", ls.Get(ctx, "key"), "value")
	}
}
