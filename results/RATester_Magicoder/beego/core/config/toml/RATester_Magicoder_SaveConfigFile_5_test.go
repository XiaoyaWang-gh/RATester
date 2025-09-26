package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestSaveConfigFile_5(t *testing.T) {
	type fields struct {
		t *toml.Tree
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &configContainer{
				t: tt.fields.t,
			}
			if err := c.SaveConfigFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("configContainer.SaveConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
