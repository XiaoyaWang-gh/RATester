package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDbQueryLog_1(t *testing.T) {
	type args struct {
		alias *alias
		db    dbQuerier
	}
	tests := []struct {
		name string
		args args
		want dbQuerier
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

			if got := newDbQueryLog(tt.args.alias, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDbQueryLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
