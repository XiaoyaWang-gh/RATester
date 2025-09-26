package ssdb

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScan_1(t *testing.T) {
	type args struct {
		keyStart string
		keyEnd   string
		limit    int
	}
	tests := []struct {
		name    string
		rc      *Cache
		args    args
		want    []string
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

			got, err := tt.rc.Scan(tt.args.keyStart, tt.args.keyEnd, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}
