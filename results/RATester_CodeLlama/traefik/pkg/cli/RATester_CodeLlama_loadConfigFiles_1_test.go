package cli

import (
	"fmt"
	"testing"
)

func TestLoadConfigFiles_1(t *testing.T) {
	type args struct {
		configFile string
		element    interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				configFile: "test case 1",
				element:    "test case 1",
			},
			want:    "test case 1",
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				configFile: "test case 2",
				element:    "test case 2",
			},
			want:    "test case 2",
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

			got, err := loadConfigFiles(tt.args.configFile, tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfigFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("loadConfigFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
