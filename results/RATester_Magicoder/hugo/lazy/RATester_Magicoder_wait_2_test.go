package lazy

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func Testwait_2(t *testing.T) {
	type fields struct {
		initCount uint64
		mu        sync.Mutex
		prev      *Init
		children  []*Init
		init      onceMore
		out       any
		err       error
		f         func(context.Context) (any, error)
	}
	tests := []struct {
		name   string
		fields fields
		want   any
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

			ini := &Init{
				initCount: tt.fields.initCount,
				mu:        tt.fields.mu,
				prev:      tt.fields.prev,
				children:  tt.fields.children,
				init:      tt.fields.init,
				out:       tt.fields.out,
				err:       tt.fields.err,
				f:         tt.fields.f,
			}
			ini.wait()
		})
	}
}
