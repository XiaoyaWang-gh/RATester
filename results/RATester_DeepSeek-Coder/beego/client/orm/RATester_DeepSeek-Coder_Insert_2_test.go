package orm

import (
	"fmt"
	"testing"
)

func TestInsert_2(t *testing.T) {
	type args struct {
		md interface{}
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

			got, err := tt.f.Insert(tt.args.md)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("filterOrmDecorator.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
