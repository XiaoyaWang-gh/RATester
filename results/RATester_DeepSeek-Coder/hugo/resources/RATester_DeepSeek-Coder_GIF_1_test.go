package resources

import (
	"fmt"
	"image/gif"
	"reflect"
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

	expected := &gif.GIF{}
	result := g.GIF()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
