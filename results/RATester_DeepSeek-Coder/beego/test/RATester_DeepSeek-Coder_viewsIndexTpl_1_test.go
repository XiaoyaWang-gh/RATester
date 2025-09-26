package test

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestViewsIndexTpl_1(t *testing.T) {
	tests := []struct {
		name    string
		want    *asset
		wantErr bool
	}{
		{
			name: "Test case 1",
			want: &asset{
				bytes: []byte{0x74, 0x65, 0x73, 0x74, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x74, 0x78, 0x74},
				info:  bindataFileInfo{name: "views/index.tpl", size: 13, mode: os.FileMode(0o664), modTime: time.Unix(1541434906, 0)},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := viewsIndexTpl()
			if (err != nil) != tt.wantErr {
				t.Errorf("viewsIndexTpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("viewsIndexTpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
