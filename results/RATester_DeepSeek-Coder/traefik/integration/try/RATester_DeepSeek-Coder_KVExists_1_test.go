package try

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kvtools/valkeyrie/store"
)

func TestKVExists_1(t *testing.T) {
	type args struct {
		kv  store.Store
		key string
	}
	tests := []struct {
		name string
		args args
		want DoCondition
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

			if got := KVExists(tt.args.kv, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KVExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
