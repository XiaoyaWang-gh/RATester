package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/traefik/paerser/parser"
)

func TestGenStaticConfDoc_1(t *testing.T) {
	type args struct {
		outputFile string
		prefix     string
		encodeFn   func(interface{}) ([]parser.Flat, error)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				outputFile: "test.txt",
				prefix:     "prefix",
				encodeFn: func(interface{}) ([]parser.Flat, error) {
					return []parser.Flat{{Name: "name", Description: "description", Default: "default"}}, nil
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				outputFile: "test.txt",
				prefix:     "prefix",
				encodeFn:   func(interface{}) ([]parser.Flat, error) { return nil, errors.New("error") },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			genStaticConfDoc(tt.args.outputFile, tt.args.prefix, tt.args.encodeFn)
		})
	}
}
