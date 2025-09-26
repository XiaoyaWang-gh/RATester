package acme

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_2(t *testing.T) {
	type args struct {
		resolverName string
	}
	tests := []struct {
		name    string
		s       *LocalStore
		args    args
		want    *StoredData
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

			got, err := tt.s.get(tt.args.resolverName)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalStore.get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalStore.get() = %v, want %v", got, tt.want)
			}
		})
	}
}
