package logs

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/rs/zerolog"
)

func TestPrintf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := new(bytes.Buffer)
	logger := zerolog.New(output)
	l := LogrusStdWrapper{logger: logger}

	testMsg := "test message"
	testArgs := "test args"
	expectedOutput := fmt.Sprintf("{\"message\":\"%s %s\"}\n", testMsg, testArgs)

	l.Printf(testMsg, testArgs)

	if got := output.String(); got != expectedOutput {
		t.Errorf("expected %q, got %q", expectedOutput, got)
	}
}
