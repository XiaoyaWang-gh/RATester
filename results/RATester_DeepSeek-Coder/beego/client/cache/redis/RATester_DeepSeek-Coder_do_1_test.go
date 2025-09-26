package redis

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDo_1(t *testing.T) {
	type args struct {
		commandName string
		args        []interface{}
	}
	tests := []struct {
		name    string
		rc      *Cache
		args    args
		want    interface{}
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

			got, err := tt.rc.do(tt.args.commandName, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.do() = %v, want %v", got, tt.want)
			}
		})
	}
}
