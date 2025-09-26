package cast

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal"
)

func Testinit_52(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := New()

	ns := &internal.TemplateFuncsNamespace{
		Name:    name,
		Context: func(cctx context.Context, args ...any) (any, error) { return ctx, nil },
	}

	ns.AddMethodMapping(ctx.ToInt,
		[]string{"int"},
		[][2]string{
			{`{{ "1234" | int | printf "%T" }}`, `int`},
		},
	)

	ns.AddMethodMapping(ctx.ToString,
		[]string{"string"},
		[][2]string{
			{`{{ 1234 | string | printf "%T" }}`, `string`},
		},
	)

	ns.AddMethodMapping(ctx.ToFloat,
		[]string{"float"},
		[][2]string{
			{`{{ "1234" | float | printf "%T" }}`, `float64`},
		},
	)

	got := ns.Name
	want := "cast"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
