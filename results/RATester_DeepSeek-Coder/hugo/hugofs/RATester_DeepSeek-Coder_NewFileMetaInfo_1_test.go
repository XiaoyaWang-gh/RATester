package hugofs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFileMetaInfo_1(t *testing.T) {
	type args struct {
		fi FileNameIsDir
		m  *FileMeta
	}
	tests := []struct {
		name string
		args args
		want FileMetaInfo
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

			if got := NewFileMetaInfo(tt.args.fi, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileMetaInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
