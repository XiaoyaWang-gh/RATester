package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestFlush_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &multiFileLogWriter{}
	f.writers = [LevelDebug + 1 + 1]*fileLogWriter{
		{
			fileWriter: &os.File{},
		},
	}
	f.Flush()
}
