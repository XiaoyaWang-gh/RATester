package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_25(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		cache  *_dbCache
		args   args
		wantAl *alias
		wantOk bool
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

			gotAl, gotOk := tt.cache.get(tt.args.name)
			if !reflect.DeepEqual(gotAl, tt.wantAl) {
				t.Errorf("get() gotAl = %v, want %v", gotAl, tt.wantAl)
			}
			if gotOk != tt.wantOk {
				t.Errorf("get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
