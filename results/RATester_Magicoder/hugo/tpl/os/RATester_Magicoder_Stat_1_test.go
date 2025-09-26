package os

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"
)

func TestStat_1(t *testing.T) {
	type args struct {
		i any
	}
	tests := []struct {
		name    string
		args    args
		want    fs.FileInfo
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

			ns := &Namespace{}
			got, err := ns.Stat(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Stat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Stat() = %v, want %v", got, tt.want)
			}
		})
	}
}
