package web

import (
	"fmt"
	"testing"
)

func TestInitBeegoBeforeTest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	InitBeegoBeforeTest(appConfigPath)
}
