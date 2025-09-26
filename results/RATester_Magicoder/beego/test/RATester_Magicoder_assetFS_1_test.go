package test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

func TestassetFS_1(t *testing.T) {
	tests := []struct {
		name string
		want *assetfs.AssetFS
	}{
		{
			name: "test1",
			want: &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: os.Stat, Prefix: "test1"},
		},
		{
			name: "test2",
			want: &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: os.Stat, Prefix: "test2"},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := assetFS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assetFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
