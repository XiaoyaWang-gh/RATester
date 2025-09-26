package memory

import (
	"fmt"
	"testing"
	"time"
)

func TestSet_2(t *testing.T) {
	type args struct {
		key string
		val []byte
		exp time.Duration
	}

	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		{
			name: "TestSet_EmptyKey",
			s:    &Storage{db: make(map[string]entry)},
			args: args{
				key: "",
				val: []byte("value"),
				exp: 0,
			},
			wantErr: false,
		},
		{
			name: "TestSet_EmptyValue",
			s:    &Storage{db: make(map[string]entry)},
			args: args{
				key: "key",
				val: []byte(""),
				exp: 0,
			},
			wantErr: false,
		},
		{
			name: "TestSet_ValidKeyValue",
			s:    &Storage{db: make(map[string]entry)},
			args: args{
				key: "key",
				val: []byte("value"),
				exp: 1 * time.Second,
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

			if err := tt.s.Set(tt.args.key, tt.args.val, tt.args.exp); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
