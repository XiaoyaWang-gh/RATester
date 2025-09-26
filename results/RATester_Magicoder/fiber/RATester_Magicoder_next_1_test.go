package fiber

import (
	"fmt"
	"testing"
)

func Testnext_1(t *testing.T) {
	app := &App{}
	ctx := &DefaultCtx{}

	tests := []struct {
		name    string
		app     *App
		ctx     *DefaultCtx
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := app.next(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}
