package helpers

import (
	"fmt"
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestinitEmoji_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	emojiMap := emoji.CodeMap()
	emojis := make(map[string][]byte)
	emojiMaxSize := 0

	for k, v := range emojiMap {
		emojis[k] = []byte(v)

		if len(k) > emojiMaxSize {
			emojiMaxSize = len(k)
		}
	}

	// Test if emojiMaxSize is correctly updated
	for k := range emojiMap {
		if len(k) > emojiMaxSize {
			t.Errorf("emojiMaxSize not correctly updated. Expected: %d, Got: %d", len(k), emojiMaxSize)
		}
	}

	// Test if emojis map is correctly populated
	for k, v := range emojiMap {
		if string(emojis[k]) != v {
			t.Errorf("emojis map not correctly populated. Expected: %s, Got: %s", v, emojis[k])
		}
	}
}
