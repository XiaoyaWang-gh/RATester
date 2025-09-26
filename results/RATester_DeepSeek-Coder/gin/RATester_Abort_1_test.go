package gin

import (
	"fmt"
	"testing"
)

func TestAbort_1(t *testing.T) {
	tests := []struct {
		name string
		ctx  *Context
		want int8
	}{
		{
			name: "TestAbort",
			ctx:  &Context{index: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.ctx.Abort()
			if got := tt.ctx.index; got != tt.want {
				t.Errorf("Abort() = %v, want %v", got, tt.want)
			}
		})
	}
}
