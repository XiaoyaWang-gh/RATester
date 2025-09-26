package xml

import (
	"fmt"
	"testing"
)

func TestInt64_10(t *testing.T) {
	t.Parallel()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "1234567890",
			"key2": "9223372036854775807",
			"key3": "9223372036854775808",
			"key4": "not_an_int",
		},
	}

	tests := []struct {
		name    string
		key     string
		want    int64
		wantErr bool
	}{
		{"valid int64", "key1", 1234567890, false},
		{"max int64", "key2", 9223372036854775807, false},
		{"overflow int64", "key3", 0, true},
		{"not an int64", "key4", 0, true},
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
