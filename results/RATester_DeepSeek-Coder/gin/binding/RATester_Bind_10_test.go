package binding

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBind_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := msgpackBinding{}
	req, _ := http.NewRequest("POST", "/", strings.NewReader(""))
	obj := make(map[string]interface{})
	err := b.Bind(req, &obj)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
