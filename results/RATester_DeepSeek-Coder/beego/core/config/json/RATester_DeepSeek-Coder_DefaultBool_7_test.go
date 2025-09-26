package json

import (
	"fmt"
	"sync"
	"testing"
)

func TestDefaultBool_7(t *testing.T) {
	type fields struct {
		data    map[string]interface{}
		RWMutex sync.RWMutex
	}
	type args struct {
		key        string
		defaultVal bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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
			if got := c.DefaultBool(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
