package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestHasReturningID_1(t *testing.T) {
	type args struct {
		mi    *models.ModelInfo
		query *string
	}
	tests := []struct {
		name string
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

			d := &dbBasePostgres{}
			if got := d.HasReturningID(tt.args.mi, tt.args.query); got != tt.want {
				t.Errorf("HasReturningID() = %v, want %v", got, tt.want)
			}
		})
	}
}
