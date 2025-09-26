package orm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestDeleteBatch_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		q    dbQuerier
		qs   *querySet
		mi   *models.ModelInfo
		cond *Condition
		tz   *time.Location
	}
	tests := []struct {
		name    string
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

			d := &dbBase{}
			got, err := d.DeleteBatch(tt.args.ctx, tt.args.q, tt.args.qs, tt.args.mi, tt.args.cond, tt.args.tz)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteBatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
