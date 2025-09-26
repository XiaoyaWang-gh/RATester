package integrity

import (
	"fmt"
	"hash"
	"reflect"
	"testing"
)

func Testdigest_1(t *testing.T) {
	type args struct {
		h hash.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
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

			got, err := digest(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("digest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digest() = %v, want %v", got, tt.want)
			}
		})
	}
}
