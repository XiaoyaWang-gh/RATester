package logs

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestFatalf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stdout)
	l := LogrusStdWrapper{logger: logger}

	msg := "test message"
	expectedOutput := fmt.Sprintf("{\"level\":\"fatal\",\"message\":\"%s\"}\n", msg)

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	l.Fatalf(msg)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(out))
	}
}
