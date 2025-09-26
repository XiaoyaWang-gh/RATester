package modules

import (
	"fmt"
	"testing"
)

func TestcreateThemeDirname_1(t *testing.T) {
	type args struct {
		modulePath   string
		isProjectMod bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				modulePath:   "test/module",
				isProjectMod: true,
			},
			want:    "test/module",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				modulePath:   "../test/module",
				isProjectMod: false,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				modulePath:   "test/module",
				isProjectMod: false,
			},
			want:    "themesDir/test/module",
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

			c := &Client{
				ccfg: ClientConfig{
					ThemesDir: "themesDir",
				},
			}
			got, err := c.createThemeDirname(tt.args.modulePath, tt.args.isProjectMod)
			if (err != nil) != tt.wantErr {
				t.Errorf("createThemeDirname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createThemeDirname() = %v, want %v", got, tt.want)
			}
		})
	}
}
