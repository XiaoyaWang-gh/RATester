package binding

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBind_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := protobufBinding{}
	req, _ := http.NewRequest("POST", "/", strings.NewReader("test"))
	obj := &struct{}{}
	err := b.Bind(req, obj)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
