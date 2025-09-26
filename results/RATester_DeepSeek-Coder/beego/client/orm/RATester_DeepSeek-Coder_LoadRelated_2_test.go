package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/utils"
)

func TestLoadRelated_2(t *testing.T) {
	type args struct {
		md   interface{}
		name string
		args []utils.KV
	}
	tests := []struct {
		name    string
		f       *filterOrmDecorator
		args    args
		want    int64
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

			got, err := tt.f.LoadRelated(tt.args.md, tt.args.name, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.LoadRelated() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("filterOrmDecorator.LoadRelated() = %v, want %v", got, tt.want)
			}
		})
	}
}
