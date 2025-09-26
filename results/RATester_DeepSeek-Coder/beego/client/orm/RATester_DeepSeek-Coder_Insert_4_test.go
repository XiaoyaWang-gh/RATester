package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestInsert_4(t *testing.T) {
	type args struct {
		ctx context.Context
		q   dbQuerier
		mi  *models.ModelInfo
		ind reflect.Value
		tz  *time.Location
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
			got, err := d.Insert(tt.args.ctx, tt.args.q, tt.args.mi, tt.args.ind, tt.args.tz)
			if (err != nil) != tt.wantErr {
				t.Errorf("dbBase.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("dbBase.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
