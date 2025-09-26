package session

import (
	"crypto/cipher"
	"fmt"
	"reflect"
	"testing"
)

func TestEncrypt_1(t *testing.T) {
	type args struct {
		block cipher.Block
		value []byte
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

			got, err := encrypt(tt.args.block, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
