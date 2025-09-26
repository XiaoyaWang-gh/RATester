package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStaticFS_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.StaticFS("/", http.Dir("."))
}
