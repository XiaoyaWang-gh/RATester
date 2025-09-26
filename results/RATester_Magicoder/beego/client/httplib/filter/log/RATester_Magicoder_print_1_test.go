package log

import (
	"errors"
	"fmt"
	"testing"
)

func Testprint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := &logInfo{
		req:  []byte("test request"),
		resp: []byte("test response"),
		err:  errors.New("test error"),
	}

	logOutput := ""
	log := func(f interface{}, v ...interface{}) {
		logOutput += fmt.Sprintf(f.(string), v...)
	}

	info.print(log)

	expectedOutput := "Request: ====================\n\"test request\"\nResponse: ===================\n\"test response\"\nError: ======================\n\"test error\"\n"
	if logOutput != expectedOutput {
		t.Errorf("Expected log output: %s, but got: %s", expectedOutput, logOutput)
	}
}
