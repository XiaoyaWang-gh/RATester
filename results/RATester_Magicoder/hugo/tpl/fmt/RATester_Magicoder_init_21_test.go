package fmt

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func Testinit_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &deps.Deps{}
	ctx := New(d)

	if ctx == nil {
		t.Error("Expected ctx to be not nil")
	}

	if ctx.Print("works") != "works" {
		t.Error("Expected Print to return 'works'")
	}

	if ctx.Println("works") != "works\n" {
		t.Error("Expected Println to return 'works\\n'")
	}

	if ctx.Printf("%s!", "works") != "works!" {
		t.Error("Expected Printf to return 'works!'")
	}

	if ctx.Errorf("%s.", "failed") != "" {
		t.Error("Expected Errorf to return ''")
	}

	if ctx.Erroridf("my-err-id", "%s.", "failed") != "" {
		t.Error("Expected Erroridf to return ''")
	}

	if ctx.Warnidf("my-warn-id", "%s.", "warning") != "" {
		t.Error("Expected Warnidf to return ''")
	}

	if ctx.Warnf("%s.", "warning") != "" {
		t.Error("Expected Warnf to return ''")
	}
}
