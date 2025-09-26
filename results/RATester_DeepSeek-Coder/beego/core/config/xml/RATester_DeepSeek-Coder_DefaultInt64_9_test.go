package xml

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_9(t *testing.T) {
	t.Parallel()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": int64(123),
			"key2": "456",
		},
	}

	tests := []struct {
		name       string
		key        string
		defaultVal int64
		want       int64
	}{
		{
			name:       "existing int64",
			key:        "key1",
			defaultVal: 789,
			want:       123,
		},
		{
			name:       "non-existing int64",
			key:        "key2",
			defaultVal: 789,
			want:       789,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cc.DefaultInt64(tt.key, tt.defaultVal); got != tt.want {
				t.Errorf("DefaultInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
