package orm

import (
	"fmt"
	"testing"
)

func TestReadOrCreate_1(t *testing.T) {
	type args struct {
		md   interface{}
		col1 string
		cols []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   int64
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
			got, got1, err := f.ReadOrCreate(tt.args.md, tt.args.col1, tt.args.cols...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadOrCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadOrCreate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReadOrCreate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
