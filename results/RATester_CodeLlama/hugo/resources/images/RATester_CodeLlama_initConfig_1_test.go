package images

import (
	"fmt"
	"testing"
)

func TestInitConfig_1(t *testing.T) {
	type fields struct {
		Format Format
		Proc   *ImageProcessor
		Spec   Spec
		*imageConfig
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"", fields{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			i := &Image{
				Format: tt.fields.Format,
				Proc:   tt.fields.Proc,
				Spec:   tt.fields.Spec,
				imageConfig: &imageConfig{
					config:       tt.fields.config,
					configInit:   tt.fields.configInit,
					configLoaded: tt.fields.configLoaded,
				},
			}
			if err := i.initConfig(); (err != nil) != tt.wantErr {
				t.Errorf("initConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
