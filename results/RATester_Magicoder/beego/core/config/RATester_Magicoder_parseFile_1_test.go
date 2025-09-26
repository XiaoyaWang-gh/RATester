package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseFile_1(t *testing.T) {
	ini := &IniConfig{}

	tests := []struct {
		name    string
		file    string
		want    *IniConfigContainer
		wantErr bool
	}{
		{
			name: "test1",
			file: "testdata/test.ini",
			want: &IniConfigContainer{
				data: map[string]map[string]string{
					"section1": {
						"key1": "value1",
						"key2": "value2",
					},
					"section2": {
						"key3": "value3",
						"key4": "value4",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "test2",
			file:    "testdata/not_exist.ini",
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

			got, err := ini.parseFile(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
