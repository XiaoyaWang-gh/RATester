package pagination

import (
	"fmt"
	"testing"
)

func TestToInt64_2(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		wantD   int64
		wantErr bool
	}{
		{"Test int", 1, 1, false},
		{"Test int8", int8(8), 8, false},
		{"Test int16", int16(16), 16, false},
		{"Test int32", int32(32), 32, false},
		{"Test int64", int64(64), 64, false},
		{"Test uint", uint(1), 1, false},
		{"Test uint8", uint8(8), 8, false},
		{"Test uint16", uint16(16), 16, false},
		{"Test uint32", uint32(32), 32, false},
		{"Test uint64", uint64(64), 64, false},
		{"Test string", "123", 0, true},
		{"Test float32", float32(32.99), 0, true},
		{"Test float64", float64(64.99), 0, true},
		{"Test bool", true, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotD, err := toInt64(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("toInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotD != tt.wantD {
				t.Errorf("toInt64() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}
