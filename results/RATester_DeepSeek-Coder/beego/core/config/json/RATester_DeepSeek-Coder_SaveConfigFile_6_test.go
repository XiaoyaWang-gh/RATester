package json

import (
	"fmt"
	"sync"
	"testing"
)

func TestSaveConfigFile_6(t *testing.T) {
	type fields struct {
		data map[string]interface{}
		sync.RWMutex
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

			c := &JSONConfigContainer{
				data:    tt.fields.data,
				RWMutex: tt.fields.RWMutex,
			}
			if err := c.SaveConfigFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("JSONConfigContainer.SaveConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
