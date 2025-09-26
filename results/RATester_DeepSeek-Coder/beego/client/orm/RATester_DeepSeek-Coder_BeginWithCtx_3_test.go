package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestBeginWithCtx_3(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		f       *filterOrmDecorator
		args    args
		want    TxOrmer
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

			got, err := tt.f.BeginWithCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.BeginWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterOrmDecorator.BeginWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
