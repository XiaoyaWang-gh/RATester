package highlight

import (
	"fmt"
	"strings"
	"testing"
)

func TestWriteDivEnd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := new(strings.Builder)
	writeDivEnd(w)
	if w.String() != "</div>" {
		t.Errorf("writeDivEnd() = %v, want %v", w.String(), "</div>")
	}
}
