package captcha

import (
	"fmt"
	"testing"
)

func TestGenRandChars_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	captcha := &Captcha{
		ChallengeNums: 6,
	}
	randChars := captcha.genRandChars()
	if len(randChars) != 6 {
		t.Errorf("Expected length of randChars is 6, but got %d", len(randChars))
	}
}
