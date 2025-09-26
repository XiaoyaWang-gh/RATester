package captcha

import (
	"fmt"
	"testing"
)

func Testkey_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	captcha := &Captcha{
		CachePrefix: "test_",
	}
	id := "123456"
	expected := "test_123456"
	result := captcha.key(id)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
