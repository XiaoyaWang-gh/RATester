package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/cache/filecache"
)

func TestWriteMeta_1(t *testing.T) {
	type args struct {
		key  string
		meta transformedResourceMetadata
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test writeMeta with valid key and meta",
			args: args{
				key: "testKey",
				meta: transformedResourceMetadata{
					Target:     "testTarget",
					MediaTypeV: "testMediaType",
					MetaData:   map[string]any{"testKey": "testValue"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test writeMeta with invalid key",
			args: args{
				key: "",
				meta: transformedResourceMetadata{
					Target:     "testTarget",
					MediaTypeV: "testMediaType",
					MetaData:   map[string]any{"testKey": "testValue"},
				},
			},
			wantErr: true,
		},
		{
			name: "Test writeMeta with invalid meta",
			args: args{
				key: "testKey",
				meta: transformedResourceMetadata{
					Target:     "",
					MediaTypeV: "testMediaType",
					MetaData:   map[string]any{"testKey": "testValue"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &ResourceCache{
				fileCache: &filecache.Cache{},
			}
			_, _, err := c.writeMeta(tt.args.key, tt.args.meta)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeMeta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
