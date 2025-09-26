package config

import (
	"fmt"
	"testing"
)

func TestSaveConfigFile_3(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		c       *fakeConfigContainer
		args    args
		wantErr bool
	}{
		{
			name: "TestSaveConfigFile_0",
			c: &fakeConfigContainer{
				data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
			args: args{
				filename: "test.json",
			},
			wantErr: true,
		},
		{
			name: "TestSaveConfigFile_1",
			c: &fakeConfigContainer{
				data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
			args: args{
				filename: "test.ini",
			},
			wantErr: true,
		},
		{
			name: "TestSaveConfigFile_2",
			c: &fakeConfigContainer{
				data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
			args: args{
				filename: "test.yaml",
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

			c := &fakeConfigContainer{
				BaseConfiger: tt.c.BaseConfiger,
				data:         tt.c.data,
			}
			if err := c.SaveConfigFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("fakeConfigContainer.SaveConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
