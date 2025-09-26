package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestDeleteWithCtx_2(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		o       querySet
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

			got, err := tt.o.DeleteWithCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("querySet.DeleteWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("querySet.DeleteWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
