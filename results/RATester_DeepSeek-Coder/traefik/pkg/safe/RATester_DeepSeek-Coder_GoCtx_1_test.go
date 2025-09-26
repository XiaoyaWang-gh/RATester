package safe

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestGoCtx_1(t *testing.T) {
	type fields struct {
		waitGroup sync.WaitGroup
		ctx       context.Context
		cancel    context.CancelFunc
	}
	type args struct {
		goroutine routineCtx
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			p := &Pool{
				waitGroup: tt.fields.waitGroup,
				ctx:       tt.fields.ctx,
				cancel:    tt.fields.cancel,
			}
			p.GoCtx(tt.args.goroutine)
		})
	}
}
