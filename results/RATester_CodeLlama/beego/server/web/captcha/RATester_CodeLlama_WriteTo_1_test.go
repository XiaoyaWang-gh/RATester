package captcha

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWriteTo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Image{}
	w := &bytes.Buffer{}
	n, err := m.WriteTo(w)
	if err != nil {
		t.Errorf("WriteTo() error = %v", err)
		return
	}
	if n != 0 {
		t.Errorf("WriteTo() n = %v, want 0", n)
	}
}
