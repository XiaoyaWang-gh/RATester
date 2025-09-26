package cache

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetMulti_1(t *testing.T) {
	type testCase struct {
		name    string
		keys    []string
		want    []interface{}
		wantErr bool
	}

	ctx := context.Background()
	cache := &MemoryCache{
		dur:   10 * time.Minute,
		items: make(map[string]*MemoryItem),
	}

	testCases := []testCase{
		{
			name:    "success",
			keys:    []string{"key1", "key2", "key3"},
			want:    []interface{}{"value1", "value2", "value3"},
			wantErr: false,
		},
		{
			name:    "failure",
			keys:    []string{"key4", "key5", "key6"},
			want:    []interface{}{"value4", "value5", "value6"},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			for i, key := range tc.keys {
				cache.Put(ctx, key, tc.want[i], cache.dur)
			}

			got, err := cache.GetMulti(ctx, tc.keys)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetMulti() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetMulti() got = %v, want %v", got, tc.want)
			}
		})
	}
}
