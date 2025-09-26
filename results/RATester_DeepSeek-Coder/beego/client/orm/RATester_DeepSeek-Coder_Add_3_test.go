package orm

import (
	"fmt"
	"testing"
)

func TestAdd_3(t *testing.T) {
	type args struct {
		mds []interface{}
	}
	tests := []struct {
		name    string
		o       *queryM2M
		args    args
		want    int64
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

			got, err := tt.o.Add(tt.args.mds...)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryM2M.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("queryM2M.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
