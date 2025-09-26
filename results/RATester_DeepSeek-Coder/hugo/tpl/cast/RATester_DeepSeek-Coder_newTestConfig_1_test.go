package cast

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTestConfig_1(t *testing.T) {
	cfg := newTestConfig()

	tests := []struct {
		name string
		key  string
		want any
	}{
		{
			name: "Test contentDir",
			key:  "contentDir",
			want: "content",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cfg.Get(tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTestConfig().Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
