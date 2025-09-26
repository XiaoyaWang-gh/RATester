package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInitColor_1(t *testing.T) {
	initColor()

	t.Run("Test colorMap", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expectedColorMap := map[string]string{
			"green":   green,
			"white":   white,
			"yellow":  yellow,
			"red":     red,
			"GET":     blue,
			"POST":    cyan,
			"PUT":     yellow,
			"DELETE":  red,
			"PATCH":   green,
			"HEAD":    magenta,
			"OPTIONS": white,
		}

		if !reflect.DeepEqual(colorMap, expectedColorMap) {
			t.Errorf("Expected colorMap to be %v, but got %v", expectedColorMap, colorMap)
		}
	})
}
