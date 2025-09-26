package resources

import (
	"fmt"
	"image/gif"
	"testing"
)

func TestGIF_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &giphy{
		gif: &gif.GIF{},
	}

	got := g.GIF()
	if got != g.gif {
		t.Errorf("g.GIF() = %v, want %v", got, g.gif)
	}
}
