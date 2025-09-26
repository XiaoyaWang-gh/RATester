package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestFlush_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{}
	w.fileWriter = &os.File{}
	w.Flush()
}
