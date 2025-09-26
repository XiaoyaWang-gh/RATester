package strings

import (
	"fmt"
	"testing"
)

func TestHasPrefix_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	tests := []struct {
		s       any
		prefix  any
		want    bool
		wantErr bool
	}{
		{"hello", "he", true, false},
		{"hello", "lo", false, false},
		{"hello", 123, false, true},
		{123, "he", false, true},
	}

	for _, tt := range tests {
		got, err := ns.HasPrefix(tt.s, tt.prefix)
		if (err != nil) != tt.wantErr {
			t.Errorf("HasPrefix() error = %v, wantErr %v", err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("HasPrefix() = %v, want %v", got, tt.want)
		}
	}
}
