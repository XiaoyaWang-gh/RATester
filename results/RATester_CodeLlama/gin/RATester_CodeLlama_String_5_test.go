package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var msgs errorMsgs
	if msgs.String() != "" {
		t.Errorf("Expected empty string, got %q", msgs.String())
	}

	msgs = append(msgs, &Error{Err: errors.New("error 1")})
	if msgs.String() != "Error #01: error 1\n" {
		t.Errorf("Expected %q, got %q", "Error #01: error 1\n", msgs.String())
	}

	msgs = append(msgs, &Error{Err: errors.New("error 2"), Meta: "meta 2"})
	if msgs.String() != "Error #01: error 1\nError #02: error 2\n     Meta: meta 2\n" {
		t.Errorf("Expected %q, got %q", "Error #01: error 1\nError #02: error 2\n     Meta: meta 2\n", msgs.String())
	}
}
