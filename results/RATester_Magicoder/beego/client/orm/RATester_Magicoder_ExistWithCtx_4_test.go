package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestExistWithCtx_4(t *testing.T) {
	ctx := context.Background()
	o := &queryM2M{
		md:  "test",
		mi:  &models.ModelInfo{},
		fi:  &models.FieldInfo{},
		qs:  &querySet{},
		ind: reflect.Value{},
	}

	tests := []struct {
		name string
		ctx  context.Context
		md   interface{}
		want bool
	}{
		{
			name: "Test case 1",
			ctx:  ctx,
			md:   "test",
			want: true,
		},
		{
			name: "Test case 2",
			ctx:  ctx,
			md:   "test2",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := o.ExistWithCtx(tt.ctx, tt.md); got != tt.want {
				t.Errorf("ExistWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
