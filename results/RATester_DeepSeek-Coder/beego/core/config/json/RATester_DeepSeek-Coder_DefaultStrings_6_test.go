package json

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestDefaultStrings_6(t *testing.T) {
	type fields struct {
		data    map[string]interface{}
		RWMutex sync.RWMutex
	}
	type args struct {
		key        string
		defaultVal []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
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
			if got := c.DefaultStrings(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
