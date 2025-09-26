package hugofs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStatRoot_1(t *testing.T) {
	type args struct {
		root     RootMapping
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    FileMetaInfo
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

			fs := &RootMappingFs{}
			got, err := fs.statRoot(tt.args.root, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("statRoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("statRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
