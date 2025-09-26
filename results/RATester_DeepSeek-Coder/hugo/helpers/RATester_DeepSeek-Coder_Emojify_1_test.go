package helpers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEmojify_1(t *testing.T) {
	tests := []struct {
		name   string
		source []byte
		want   []byte
	}{
		{
			name:   "Test case 1",
			source: []byte("Hello :smile:"),
			want:   []byte("Hello 😄"),
		},
		{
			name:   "Test case 2",
			source: []byte(":smile: :+1:"),
			want:   []byte("😄 👍"),
		},
		{
			name:   "Test case 3",
			source: []byte(":smile::+1:"),
			want:   []byte("😄:+1:"),
		},
		{
			name:   "Test case 4",
			source: []byte(":smile: :+1: :smile:"),
			want:   []byte("😄 👍 😄"),
		},
		{
			name:   "Test case 5",
			source: []byte(":smile: :+1: :smile: :+1:"),
			want:   []byte("😄 👍 😄 👍"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Emojify(tt.source); !bytes.Equal(got, tt.want) {
				t.Errorf("Emojify() = %v, want %v", got, tt.want)
			}
		})
	}
}
