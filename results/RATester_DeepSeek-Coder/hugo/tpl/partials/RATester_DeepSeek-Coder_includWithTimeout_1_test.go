package partials

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestIncludWithTimeout_1(t *testing.T) {
	type fields struct {
		deps           *deps.Deps
		cachedPartials *partialCache
	}
	type args struct {
		ctx      context.Context
		name     string
		dataList []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   includeResult
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

			ns := &Namespace{
				deps:           tt.fields.deps,
				cachedPartials: tt.fields.cachedPartials,
			}
			if got := ns.includWithTimeout(tt.args.ctx, tt.args.name, tt.args.dataList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.includWithTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
