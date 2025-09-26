package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewCtx_1(t *testing.T) {
	app := &App{}

	tests := []struct {
		name string
		app  *App
		want Ctx
	}{
		{
			name: "Test newCtx",
			app:  app,
			want: &DefaultCtx{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.app.newCtx(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
