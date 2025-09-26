package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func Testget_25(t *testing.T) {
	ac := &_dbCache{
		cache: map[string]*alias{
			"test": {
				Name: "test",
			},
		},
	}

	tests := []struct {
		name     string
		cache    *_dbCache
		arg      string
		wantAl   *alias
		wantOk   bool
		wantName string
	}{
		{
			name:     "found",
			cache:    ac,
			arg:      "test",
			wantAl:   ac.cache["test"],
			wantOk:   true,
			wantName: "test",
		},
		{
			name:     "not found",
			cache:    ac,
			arg:      "not_exist",
			wantAl:   nil,
			wantOk:   false,
			wantName: "not_exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotAl, gotOk := tt.cache.get(tt.arg)
			if !reflect.DeepEqual(gotAl, tt.wantAl) {
				t.Errorf("get() gotAl = %v, want %v", gotAl, tt.wantAl)
			}
			if gotOk != tt.wantOk {
				t.Errorf("get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
