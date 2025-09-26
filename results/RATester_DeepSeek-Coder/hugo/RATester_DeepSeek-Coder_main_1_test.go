package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain_1(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	testCases := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Test case 1",
			args:     []string{"command1", "arg1", "arg2"},
			expected: "",
		},
		{
			name:     "Test case 2",
			args:     []string{"command2", "arg1", "arg2"},
			expected: "",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Args = append([]string{"program"}, tc.args...)
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			main()

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = rescueStdout

			if strings.TrimSpace(string(out)) != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, string(out))
			}
		})
	}
}
