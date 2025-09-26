package kv

import (
	"fmt"
	"testing"

	"github.com/kvtools/valkeyrie/store"
)

func TestDecode_2(t *testing.T) {
	type args struct {
		pairs    []*store.KVPair
		element  interface{}
		rootName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				pairs: []*store.KVPair{
					{Key: "key1", Value: []byte("value1")},
					{Key: "key2", Value: []byte("value2")},
				},
				element:  &struct{}{},
				rootName: "root",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				pairs: []*store.KVPair{
					{Key: "key1", Value: []byte("value1")},
					{Key: "key2", Value: []byte("value2")},
				},
				element:  nil,
				rootName: "root",
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				pairs: []*store.KVPair{
					{Key: "key1", Value: []byte("value1")},
					{Key: "key2", Value: []byte("value2")},
				},
				element:  &struct{}{},
				rootName: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := Decode(tt.args.pairs, tt.args.element, tt.args.rootName); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
