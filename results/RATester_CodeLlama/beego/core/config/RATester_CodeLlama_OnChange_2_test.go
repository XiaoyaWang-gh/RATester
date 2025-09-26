package config

import (
	"fmt"
	"testing"
	"time"
)

func TestOnChange_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "key"
	fn := func(value string) {
		t.Log("value:", value)
	}
	OnChange(key, fn)
	value := "value"
	globalInstance.Set(key, value)
	time.Sleep(time.Second)
}
