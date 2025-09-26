package utils

import (
	"fmt"
	"sync"
	"testing"
)

func TestSet_41(t *testing.T) {
	tests := []struct {
		name string
		bm   *BeeMap
		k    interface{}
		v    interface{}
		want bool
	}{
		{
			name: "TestSet_NewKey",
			bm:   &BeeMap{lock: &sync.RWMutex{}, bm: make(map[interface{}]interface{})},
			k:    "key1",
			v:    "value1",
			want: true,
		},
		{
			name: "TestSet_ExistingKey",
			bm:   &BeeMap{lock: &sync.RWMutex{}, bm: map[interface{}]interface{}{"key1": "value1"}},
			k:    "key1",
			v:    "value1",
			want: false,
		},
		{
			name: "TestSet_DifferentValue",
			bm:   &BeeMap{lock: &sync.RWMutex{}, bm: map[interface{}]interface{}{"key1": "value1"}},
			k:    "key1",
			v:    "value2",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.bm.Set(tt.k, tt.v); got != tt.want {
				t.Errorf("BeeMap.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
