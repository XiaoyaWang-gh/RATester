package session

import (
	"crypto/cipher"
	"fmt"
	"testing"
)

func TestencodeCookie_1(t *testing.T) {
	type args struct {
		block   cipher.Block
		hashKey string
		name    string
		value   map[interface{}]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
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

			got, err := encodeCookie(tt.args.block, tt.args.hashKey, tt.args.name, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("encodeCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
