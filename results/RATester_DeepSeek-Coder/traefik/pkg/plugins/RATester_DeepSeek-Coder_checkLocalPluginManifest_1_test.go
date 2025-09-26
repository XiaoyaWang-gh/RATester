package plugins

import (
	"fmt"
	"testing"
)

func TestCheckLocalPluginManifest_1(t *testing.T) {
	type args struct {
		descriptor LocalDescriptor
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				descriptor: LocalDescriptor{
					ModuleName: "testModule",
				},
			},
			wantErr: true,
		},
		{
			name: "Test case 2",
			args: args{
				descriptor: LocalDescriptor{
					ModuleName: "testModule",
					Settings:   Settings{},
				},
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

			if err := checkLocalPluginManifest(tt.args.descriptor); (err != nil) != tt.wantErr {
				t.Errorf("checkLocalPluginManifest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
