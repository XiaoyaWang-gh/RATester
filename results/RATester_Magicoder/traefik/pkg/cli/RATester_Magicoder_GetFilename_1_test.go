package cli

import (
	"fmt"
	"testing"
)

func TestGetFilename_1(t *testing.T) {
	type fields struct {
		ConfigFileFlag string
		filename       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				ConfigFileFlag: "test.toml",
				filename:       "test.toml",
			},
			want: "test.toml",
		},
		{
			name: "Test case 2",
			fields: fields{
				ConfigFileFlag: "test2.toml",
				filename:       "test2.toml",
			},
			want: "test2.toml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &FileLoader{
				ConfigFileFlag: tt.fields.ConfigFileFlag,
				filename:       tt.fields.filename,
			}
			if got := f.GetFilename(); got != tt.want {
				t.Errorf("FileLoader.GetFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
