package session

import (
	"crypto/cipher"
	"fmt"
	"reflect"
	"testing"
)

func TestdecodeCookie_1(t *testing.T) {
	type args struct {
		block         cipher.Block
		hashKey       string
		name          string
		value         string
		gcmaxlifetime int64
	}
	tests := []struct {
		name    string
		args    args
		want    map[interface{}]interface{}
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

			got, err := decodeCookie(tt.args.block, tt.args.hashKey, tt.args.name, tt.args.value, tt.args.gcmaxlifetime)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
