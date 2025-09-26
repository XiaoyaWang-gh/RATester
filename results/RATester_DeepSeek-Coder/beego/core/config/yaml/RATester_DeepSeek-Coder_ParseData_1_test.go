package yaml

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/config"
)

func TestParseData_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		data    []byte
		want    config.Configer
		wantErr bool
	}{
		{
			name: "valid yaml",
			data: []byte(`
field: value
`),
			want: &ConfigContainer{
				data: map[string]interface{}{
					"field": "value",
				},
			},
			wantErr: false,
		},
		{
			name:    "invalid yaml",
			data:    []byte(`invalid yaml`),
			want:    nil,
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

			c := &Config{}
			got, err := c.ParseData(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
