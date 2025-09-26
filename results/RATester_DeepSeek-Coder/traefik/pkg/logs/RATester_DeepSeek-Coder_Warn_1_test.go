package logs

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestWarn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &bytes.Buffer{}
	logger := zerolog.New(output).With().Caller().Logger()
	l := OxyWrapper{logger: logger}

	testMessage := "test message"
	testArgs := []interface{}{"arg1", "arg2"}
	expectedOutput := fmt.Sprintf("{\"level\":\"warn\",\"message\":\"%s %s\",\"caller\":\"%s:%d\"}", testMessage, testArgs, "oxy_wrapper_test.go", 28)

	l.Warn(testMessage, testArgs...)

	gotOutput := strings.TrimSpace(output.String())
	if gotOutput != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, gotOutput)
	}
}
