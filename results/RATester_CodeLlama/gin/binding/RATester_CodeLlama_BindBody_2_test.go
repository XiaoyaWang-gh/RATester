package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var body []byte
	var obj any
	err := (&plainBinding{}).BindBody(body, obj)
	if err != nil {
		t.Errorf("BindBody() error = %v, want nil", err)
		return
	}
}
