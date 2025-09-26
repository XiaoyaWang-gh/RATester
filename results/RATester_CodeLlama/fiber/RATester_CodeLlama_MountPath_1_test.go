package fiber

import (
	"fmt"
	"testing"
)

func TestMountPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	app.MountPath()
	// TODO
}
