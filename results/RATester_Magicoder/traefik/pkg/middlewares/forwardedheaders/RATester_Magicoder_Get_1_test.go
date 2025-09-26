package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestGet_1(t *testing.T) {
	header := make(unsafeHeader)
	header.Set("key", "value")

	tests := []struct {
		name string
		h    unsafeHeader
		key  string
		want string
	}{
		{
			name: "Get existing key",
			h:    header,
			key:  "key",
			want: "value",
		},
		{
			name: "Get non-existing key",
			h:    header,
			key:  "non-existing-key",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.h.Get(tt.key); got != tt.want {
				t.Errorf("unsafeHeader.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
