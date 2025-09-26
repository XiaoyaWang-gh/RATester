package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := fileWriter{
		baseDir: "testdata",
	}

	files := map[string]*File{
		"testdata/test.go": {
			Package: "test",
			Imports: []string{"fmt"},
			Elements: []Element{
				{
					Name:  "test",
					Value: "test",
				},
			},
		},
	}

	err := f.Write(files)
	require.NoError(t, err)

	testdata, err := ioutil.ReadFile("testdata/test.go")
	require.NoError(t, err)

	expected := `package test

import "fmt"

var test = "test"
`
	require.Equal(t, expected, string(testdata))
}
