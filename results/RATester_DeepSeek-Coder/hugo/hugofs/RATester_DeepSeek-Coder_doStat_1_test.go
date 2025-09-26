package hugofs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDoStat_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fs      *RootMappingFs
		args    args
		want    []FileMetaInfo
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

			got, err := tt.fs.doStat(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("RootMappingFs.doStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RootMappingFs.doStat() = %v, want %v", got, tt.want)
			}
		})
	}
}
