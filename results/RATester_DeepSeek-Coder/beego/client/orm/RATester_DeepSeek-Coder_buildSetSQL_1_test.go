package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/buffers"
)

func TestBuildSetSQL_1(t *testing.T) {
	type args struct {
		buf    buffers.Buffer
		cols   []string
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
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

			d := &dbBase{}
			d.buildSetSQL(tt.args.buf, tt.args.cols, tt.args.values)
			if got := tt.args.buf.String(); got != tt.want {
				t.Errorf("buildSetSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
