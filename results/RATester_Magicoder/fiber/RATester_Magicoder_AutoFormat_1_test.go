package fiber

import (
	"fmt"
	"testing"
)

func TestAutoFormat_1(t *testing.T) {
	ctx := &DefaultCtx{}

	tests := []struct {
		name    string
		ctx     *DefaultCtx
		body    any
		wantErr bool
	}{
		{
			name: "Test case 1",
			ctx:  ctx,
			body: "test",
		},
		{
			name: "Test case 2",
			ctx:  ctx,
			body: []byte("test"),
		},
		{
			name: "Test case 3",
			ctx:  ctx,
			body: struct{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.ctx.AutoFormat(tt.body); (err != nil) != tt.wantErr {
				t.Errorf("AutoFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
