package highlight

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWritePreStart_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var w bytes.Buffer
	WritePreStart(&w, "go", "style")
	if w.String() != `<pre tabindex="0" style="">` {
		t.Errorf("WritePreStart() = %v, want %v", w.String(), `<pre tabindex="0" style="">`)
	}
}
