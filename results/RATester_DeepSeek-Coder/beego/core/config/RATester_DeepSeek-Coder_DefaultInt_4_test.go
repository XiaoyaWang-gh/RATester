package config

import (
	"fmt"
	"sync"
	"testing"
)

func TestDefaultInt_4(t *testing.T) {
	t.Parallel()

	type fields struct {
		BaseConfiger   BaseConfiger
		data           map[string]map[string]string
		sectionComment map[string]string
		keyComment     map[string]string
		RWMutex        sync.RWMutex
	}
	type args struct {
		key        string
		defaultVal int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
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
			if got := c.DefaultInt(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
