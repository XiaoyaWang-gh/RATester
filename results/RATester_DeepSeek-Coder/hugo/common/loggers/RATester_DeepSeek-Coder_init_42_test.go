package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestInit_42(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	InitGlobalLogger(logg.LevelWarn, false)
}
