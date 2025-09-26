package hugofs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestApplyMeta_1(t *testing.T) {
	type args struct {
		fi   FileNameIsDir
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  FileMetaInfo
		want1 bool
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

			fs := &componentFs{}
			got, got1 := fs.applyMeta(tt.args.fi, tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyMeta() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("applyMeta() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
