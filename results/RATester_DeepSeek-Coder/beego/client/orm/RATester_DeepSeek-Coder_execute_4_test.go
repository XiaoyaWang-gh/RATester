package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestExecute_4(t *testing.T) {
	type fields struct {
		Method      string
		Md          interface{}
		Args        []interface{}
		mi          *models.ModelInfo
		f           func(ctx context.Context) []interface{}
		InsideTx    bool
		TxStartTime time.Time
		TxName      string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
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

			inv := &Invocation{
				Method:      tt.fields.Method,
				Md:          tt.fields.Md,
				Args:        tt.fields.Args,
				mi:          tt.fields.mi,
				f:           tt.fields.f,
				InsideTx:    tt.fields.InsideTx,
				TxStartTime: tt.fields.TxStartTime,
				TxName:      tt.fields.TxName,
			}
			if got := inv.execute(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invocation.execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
