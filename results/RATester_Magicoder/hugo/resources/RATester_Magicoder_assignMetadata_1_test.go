package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestassignMetadata_1(t *testing.T) {
	type args struct {
		metadata []map[string]any
		ma       *metaResource
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				metadata: []map[string]any{
					{
						"src":   "test_src",
						"name":  "test_name",
						"title": "test_title",
						"params": map[string]any{
							"param1": "value1",
							"param2": "value2",
						},
					},
				},
				ma: &metaResource{
					name:   "initial_name",
					title:  "initial_title",
					params: maps.Params{"initial_param": "value"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				metadata: []map[string]any{
					{
						"src":   "test_src",
						"name":  "test_name",
						"title": "test_title",
						"params": map[string]any{
							"param1": "value1",
							"param2": "value2",
						},
					},
				},
				ma: &metaResource{
					name:   "initial_name",
					title:  "initial_title",
					params: maps.Params{"initial_param": "value"},
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

			if err := assignMetadata(tt.args.metadata, tt.args.ma); (err != nil) != tt.wantErr {
				t.Errorf("assignMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
