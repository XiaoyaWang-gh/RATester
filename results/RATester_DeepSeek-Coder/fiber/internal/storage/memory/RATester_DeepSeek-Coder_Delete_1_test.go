package memory

import (
	"fmt"
	"testing"
)

func TestDelete_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		{
			name: "TestDelete_0",
			s: &Storage{
				db: map[string]entry{
					"key1": {data: []byte("value1"), expiry: 0},
					"key2": {data: []byte("value2"), expiry: 0},
				},
			},
			args:    args{key: "key1"},
			wantErr: false,
		},
		{
			name: "TestDelete_1",
			s: &Storage{
				db: map[string]entry{
					"key1": {data: []byte("value1"), expiry: 0},
					"key2": {data: []byte("value2"), expiry: 0},
				},
			},
			args:    args{key: "key3"},
			wantErr: false,
		},
		{
			name: "TestDelete_2",
			s: &Storage{
				db: map[string]entry{
					"key1": {data: []byte("value1"), expiry: 0},
					"key2": {data: []byte("value2"), expiry: 0},
				},
			},
			args:    args{key: ""},
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

			if err := tt.s.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
