package orm

import (
	"fmt"
	"testing"
)

func TestDelete_17(t *testing.T) {
	type args struct {
		md   interface{}
		cols []string
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

			f := &filterOrmDecorator{}
			got, err := f.Delete(tt.args.md, tt.args.cols...)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("filterOrmDecorator.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
