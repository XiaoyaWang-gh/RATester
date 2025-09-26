package kv

import (
	"fmt"
	"testing"

	"github.com/kvtools/valkeyrie/store"
)

func TestDecode_2(t *testing.T) {
	type testStruct struct {
		Name string `kv:"name"`
		Age  int    `kv:"age"`
	}

	tests := []struct {
		name     string
		pairs    []*store.KVPair
		element  interface{}
		rootName string
		wantErr  bool
	}{
		{
			name: "success",
			pairs: []*store.KVPair{
				{Key: "name", Value: []byte("John"), LastIndex: 1},
				{Key: "age", Value: []byte("30"), LastIndex: 1},
			},
			element:  &testStruct{},
			rootName: "root",
			wantErr:  false,
		},
		{
			name: "failure",
			pairs: []*store.KVPair{
				{Key: "name", Value: []byte("John"), LastIndex: 1},
				{Key: "age", Value: []byte("not_an_int"), LastIndex: 1},
			},
			element:  &testStruct{},
			rootName: "root",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := Decode(tt.pairs, tt.element, tt.rootName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
