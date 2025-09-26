package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMessages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Redirect{
		c: &DefaultCtx{
			flashMessages: redirectionMsgs{
				{
					key:        "key1",
					value:      "value1",
					level:      1,
					isOldInput: false,
				},
				{
					key:        "key2",
					value:      "value2",
					level:      2,
					isOldInput: true,
				},
			},
		},
	}

	expected := []FlashMessage{
		{
			Key:   "key1",
			Value: "value1",
			Level: 1,
		},
	}

	result := r.Messages()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
