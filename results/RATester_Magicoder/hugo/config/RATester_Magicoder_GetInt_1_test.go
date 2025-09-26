package config

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetInt_1(t *testing.T) {
	type fields struct {
		mu       sync.RWMutex
		root     maps.Params
		keyCache sync.Map
	}
	type args struct {
		k string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
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

			c := &defaultConfigProvider{
				mu:       tt.fields.mu,
				root:     tt.fields.root,
				keyCache: tt.fields.keyCache,
			}
			if got := c.GetInt(tt.args.k); got != tt.want {
				t.Errorf("defaultConfigProvider.GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
