package config

import (
	"fmt"
	"sync"
	"testing"
)

func TestInt_3(t *testing.T) {
	type fields struct {
		BaseConfiger   BaseConfiger
		data           map[string]map[string]string
		sectionComment map[string]string
		keyComment     map[string]string
		RWMutex        sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
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

			c := &IniConfigContainer{
				BaseConfiger:   tt.fields.BaseConfiger,
				data:           tt.fields.data,
				sectionComment: tt.fields.sectionComment,
				keyComment:     tt.fields.keyComment,
				RWMutex:        tt.fields.RWMutex,
			}
			got, err := c.Int(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("IniConfigContainer.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IniConfigContainer.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
