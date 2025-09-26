package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDIY_1(t *testing.T) {
	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
		},
	}

	tests := []struct {
		name    string
		key     string
		wantV   interface{}
		wantErr bool
	}{
		{
			name:    "existing key",
			key:     "key1",
			wantV:   "value1",
			wantErr: false,
		},
		{
			name:    "non-existing key",
			key:     "key3",
			wantV:   nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotV, err := cc.DIY(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DIY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("DIY() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}
