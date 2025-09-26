package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStaticFileFS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.StaticFileFS("relativePath", "filepath", http.Dir("."))
}
