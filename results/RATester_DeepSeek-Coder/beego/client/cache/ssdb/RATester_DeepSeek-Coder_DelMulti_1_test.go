package ssdb

import (
	"fmt"
	"testing"
)

func TestDelMulti_1(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		rc      *Cache
		args    args
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

			if err := tt.rc.DelMulti(tt.args.keys); (err != nil) != tt.wantErr {
				t.Errorf("DelMulti() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
