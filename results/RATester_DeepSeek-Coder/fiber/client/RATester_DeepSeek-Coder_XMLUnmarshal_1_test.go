package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/utils/v2"
)

func TestXMLUnmarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	client := &Client{
		xmlUnmarshal: func(data []byte, v any) error {
			return nil
		},
	}

	expected := utils.XMLUnmarshal(func(data []byte, v any) error {
		return nil
	})

	result := client.XMLUnmarshal()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
