package logger

import (
	"fmt"
	"io"
	"testing"
)

func TestWriteLog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var w io.Writer
	var msg []byte
	writeLog(w, msg)
}
