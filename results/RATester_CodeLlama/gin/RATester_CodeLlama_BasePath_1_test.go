package gin

import (
	"fmt"
	"testing"
)

func TestBasePath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	group := &RouterGroup{
		basePath: "/base",
	}
	if group.BasePath() != "/base" {
		t.Errorf("BasePath() = %v, want %v", group.BasePath(), "/base")
	}
}
