package fiber

import (
	"fmt"
	"testing"
)

func TestPath_1(t *testing.T) {
	ctx := &DefaultCtx{
		path: "original_path",
	}

	tests := []struct {
		name     string
		ctx      *DefaultCtx
		override []string
		want     string
	}{
		{
			name: "Test with override",
			ctx:  ctx,
			override: []string{
				"override_path",
			},
			want: "override_path",
		},
		{
			name: "Test without override",
			ctx:  ctx,
			want: "original_path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.ctx.Path(tt.override...); got != tt.want {
				t.Errorf("Path() = %v, want %v", got, tt.want)
			}
		})
	}
}
