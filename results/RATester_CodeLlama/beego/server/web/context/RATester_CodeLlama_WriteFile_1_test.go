package context

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestWriteFile_1(t *testing.T) {
	t.Run("TestWriteFile", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var encoding string
		var writer io.Writer
		var file *os.File
		var want bool
		var want1 string
		var want2 error
		got, got1, got2 := WriteFile(encoding, writer, file)
		if (got != want) || (got1 != want1) || (got2 != want2) {
			t.Errorf("WriteFile() = %v, %v, %v, want %v, %v, %v", got, got1, got2, want, want1, want2)
		}
	})
}
