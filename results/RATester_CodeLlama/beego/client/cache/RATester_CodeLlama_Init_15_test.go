package cache

import (
	"fmt"
	"testing"
)

func TestInit_15(t *testing.T) {
	type fields struct {
		CachePath      string
		FileSuffix     string
		DirectoryLevel int
		EmbedExpiry    int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "TestInit",
			fields: fields{
				CachePath:      "./cache",
				FileSuffix:     ".cache",
				DirectoryLevel: 2,
				EmbedExpiry:    0,
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

			fc := &FileCache{
				CachePath:      tt.fields.CachePath,
				FileSuffix:     tt.fields.FileSuffix,
				DirectoryLevel: tt.fields.DirectoryLevel,
				EmbedExpiry:    tt.fields.EmbedExpiry,
			}
			if err := fc.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
