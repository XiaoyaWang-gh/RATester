package related

import (
	"fmt"
	"testing"
)

func TestHasType_1(t *testing.T) {
	cfg := &Config{
		Threshold:    50,
		IncludeNewer: true,
		ToLower:      true,
		Indices: IndicesConfig{
			{Type: "type1"},
			{Type: "type2"},
			{Type: "type3"},
		},
	}

	tests := []struct {
		name string
		cfg  *Config
		arg  string
		want bool
	}{
		{
			name: "has type",
			cfg:  cfg,
			arg:  "type2",
			want: true,
		},
		{
			name: "has not type",
			cfg:  cfg,
			arg:  "type4",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.cfg.HasType(tt.arg)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
