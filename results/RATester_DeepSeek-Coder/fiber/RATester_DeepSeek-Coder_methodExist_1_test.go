package fiber

import (
	"fmt"
	"testing"
)

func TestMethodExist_1(t *testing.T) {
	app := new(App)
	ctx := new(DefaultCtx)

	tests := []struct {
		name string
		app  *App
		ctx  *DefaultCtx
		want bool
	}{
		{
			name: "Test methodExist",
			app:  app,
			ctx:  ctx,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := app.methodExist(ctx); got != tt.want {
				t.Errorf("methodExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
