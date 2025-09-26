package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestHasReturningID_2(t *testing.T) {
	type args struct {
		mi  *models.ModelInfo
		col *string
	}
	tests := []struct {
		name string
		d    *dbBase
		args args
		want bool
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

			if got := tt.d.HasReturningID(tt.args.mi, tt.args.col); got != tt.want {
				t.Errorf("dbBase.HasReturningID() = %v, want %v", got, tt.want)
			}
		})
	}
}
