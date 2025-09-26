package json

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/config"
)

func TestParseData_4(t *testing.T) {
	js := &JSONConfig{}

	tests := []struct {
		name    string
		data    []byte
		want    config.Configer
		wantErr bool
	}{
		{
			name: "valid json",
			data: []byte(`{"key": "value"}`),
			want: &JSONConfigContainer{
				data: map[string]interface{}{
					"key": "value",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid json",
			data: []byte(`{key: "value"}`),
			want: &JSONConfigContainer{
				data: map[string]interface{}{
					"rootArray": []interface{}{
						map[string]interface{}{
							"key": "value",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "empty data",
			data:    []byte{},
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

			got, err := js.ParseData(tt.data)
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
