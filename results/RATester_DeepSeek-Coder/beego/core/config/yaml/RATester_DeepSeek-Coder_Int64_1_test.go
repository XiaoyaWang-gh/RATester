package yaml

import (
	"fmt"
	"testing"
)

func TestInt64_1(t *testing.T) {
	t.Parallel()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": 123,
			"key2": int64(456),
			"key3": "not int or int64 value",
		},
	}

	tests := []struct {
		name    string
		key     string
		want    int64
		wantErr bool
	}{
		{"int", "key1", 123, false},
		{"int64", "key2", 456, false},
		{"not int or int64 value", "key3", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := cc.Int64(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}
