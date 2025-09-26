package logs

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &bytes.Buffer{}
	logger := zerolog.New(output).With().Timestamp().Logger()
	l := OxyWrapper{logger: logger}

	testString := "test error"
	testArgs := []interface{}{"arg1", "arg2"}
	expectedOutput := fmt.Sprintf("{\"level\":\"error\",\"time\":%s,\"message\":\"%s %s %s\"}\n", "%s", testString, testArgs[0], testArgs[1])

	l.Error(testString, testArgs...)

	gotOutput := output.String()
	if !strings.Contains(gotOutput, expectedOutput) {
		t.Errorf("Expected output to contain %s, got %s", expectedOutput, gotOutput)
	}
}
