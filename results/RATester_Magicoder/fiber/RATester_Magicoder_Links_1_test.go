package fiber

import (
	"fmt"
	"testing"
)

func TestLinks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	ctx := app.AcquireCtx(nil)
	defer app.ReleaseCtx(ctx)

	ctx.Links("<https://example.com>", "next", "<https://example.com/next>", "prev", "<https://example.com/prev>")

	expected := `<https://example.com>; rel="next",<https://example.com/next>; rel="prev",`
	if ctx.GetRespHeader(HeaderLink) != expected {
		t.Errorf("Expected %s, got %s", expected, ctx.GetRespHeader(HeaderLink))
	}
}
