package web

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGetMapData_1(t *testing.T) {
	type fields struct {
		lock        sync.RWMutex
		LengthLimit int
		urlmap      map[string]map[string]*Statistics
	}
	tests := []struct {
		name   string
		fields fields
		want   []map[string]interface{}
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

			m := &URLMap{
				lock:        tt.fields.lock,
				LengthLimit: tt.fields.LengthLimit,
				urlmap:      tt.fields.urlmap,
			}
			if got := m.GetMapData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLMap.GetMapData() = %v, want %v", got, tt.want)
			}
		})
	}
}
