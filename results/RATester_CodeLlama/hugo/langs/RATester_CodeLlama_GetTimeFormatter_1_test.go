package langs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetTimeFormatter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	l := &Language{}
	// When
	timeFormatter := GetTimeFormatter(l)
	// Then
	assert.NotNil(t, timeFormatter)
}
