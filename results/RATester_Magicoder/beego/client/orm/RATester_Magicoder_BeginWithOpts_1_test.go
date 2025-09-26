package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestBeginWithOpts_1(t *testing.T) {
	type args struct {
		opts *sql.TxOptions
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

			got, err := tt.f.BeginWithOpts(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.BeginWithOpts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterOrmDecorator.BeginWithOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}
