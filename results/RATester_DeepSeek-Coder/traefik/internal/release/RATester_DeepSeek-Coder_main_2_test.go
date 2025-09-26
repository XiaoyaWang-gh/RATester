package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain_2(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	cases := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Test with GOOS provided",
			args:     []string{"program", "linux"},
			expected: "linux",
		},
		{
			name:     "Test with no GOOS provided",
			args:     []string{"program"},
			expected: "GOOS should be provided as a CLI argument",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Args = tc.args
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			main()

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = rescueStdout

			if strings.TrimSpace(string(out)) != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, string(out))
			}
		})
	}
}
