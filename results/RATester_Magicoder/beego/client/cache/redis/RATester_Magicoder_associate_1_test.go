package redis

import (
	"fmt"
	"testing"
)

func Testassociate_1(t *testing.T) {
	rc := &Cache{
		key:             "test",
		skipEmptyPrefix: true,
	}

	tests := []struct {
		name       string
		originKey  interface{}
		wantResult string
	}{
		{
			name:       "Test with string",
			originKey:  "testKey",
			wantResult: "test:testKey",
		},
		{
			name:       "Test with int",
			originKey:  123,
			wantResult: "test:123",
		},
		{
			name:       "Test with empty key",
			originKey:  "",
			wantResult: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := rc.associate(tt.originKey); got != tt.wantResult {
				t.Errorf("associate() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}
