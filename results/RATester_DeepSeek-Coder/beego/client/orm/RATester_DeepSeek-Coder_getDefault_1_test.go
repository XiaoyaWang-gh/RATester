package orm

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGetDefault_1(t *testing.T) {
	type fields struct {
		mux   sync.RWMutex
		cache map[string]*alias
	}
	tests := []struct {
		name   string
		fields fields
		wantAl *alias
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

			ac := &_dbCache{
				mux:   tt.fields.mux,
				cache: tt.fields.cache,
			}
			if gotAl := ac.getDefault(); !reflect.DeepEqual(gotAl, tt.wantAl) {
				t.Errorf("_dbCache.getDefault() = %v, want %v", gotAl, tt.wantAl)
			}
		})
	}
}
