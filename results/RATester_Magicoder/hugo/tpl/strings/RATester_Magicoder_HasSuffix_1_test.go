package strings

import (
	"fmt"
	"testing"
)

func TestHasSuffix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	tests := []struct {
		s       any
		suffix  any
		want    bool
		wantErr bool
	}{
		{"hello", "lo", true, false},
		{"hello", "world", false, false},
		{"hello", 123, false, true},
		{123, "lo", false, true},
	}

	for _, tt := range tests {
		got, err := ns.HasSuffix(tt.s, tt.suffix)
		if (err != nil) != tt.wantErr {
			t.Errorf("HasSuffix() error = %v, wantErr %v", err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("HasSuffix() = %v, want %v", got, tt.want)
		}
	}
}
