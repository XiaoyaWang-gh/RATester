package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestDestroy_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{}
	w.fileWriter = &os.File{}
	w.Destroy()
	if w.fileWriter != nil {
		t.Error("file writer should be nil")
	}
}
