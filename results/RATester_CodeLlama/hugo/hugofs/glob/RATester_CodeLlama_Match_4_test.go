package glob

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestMatch_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := globDecorator{
		isWindows: true,
		g:         glob.MustCompile("*"),
	}
	if !g.Match("C:\\foo\\bar") {
		t.Errorf("expected true, got false")
	}
}
