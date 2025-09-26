package helpers

import (
	"fmt"
	"testing"
)

func TestInitEmoji_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	emojis = make(map[string][]byte)
	emojiMaxSize = 0

	initEmoji()

	if len(emojis) == 0 {
		t.Errorf("Expected emojis map to be populated, but it's empty")
	}

	if emojiMaxSize == 0 {
		t.Errorf("Expected emojiMaxSize to be greater than 0, but it's 0")
	}

	for k, v := range emojis {
		if len(k) > emojiMaxSize {
			t.Errorf("Expected emoji key '%s' to be less than or equal to emojiMaxSize, but it's greater", k)
		}

		if len(v) == 0 {
			t.Errorf("Expected emoji value for key '%s' to be populated, but it's empty", k)
		}
	}
}
