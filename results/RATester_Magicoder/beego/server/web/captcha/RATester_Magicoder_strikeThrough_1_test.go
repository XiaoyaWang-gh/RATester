package captcha

import (
	"fmt"
	"testing"
)

func TeststrikeThrough_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	img := &Image{
		numWidth:  10,
		numHeight: 10,
		dotSize:   1,
	}
	img.strikeThrough()
	// Add assertions here
}
