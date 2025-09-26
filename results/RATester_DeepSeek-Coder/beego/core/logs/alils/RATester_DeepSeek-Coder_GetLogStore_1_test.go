package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLogStore_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		p       *LogProject
		args    args
		wantS   *LogStore
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

			gotS, err := tt.p.GetLogStore(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLogStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("GetLogStore() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
