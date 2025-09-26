package lazy

import (
	"fmt"
	"testing"
)

func TestDo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	once := &onceMore{}
	once.Do(func() {
		fmt.Println("Hello")
	})
	once.Do(func() {
		fmt.Println("Hello")
	})
}
