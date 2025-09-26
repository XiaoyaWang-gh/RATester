package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGetBuffer_1(t *testing.T) {
	type args struct {
		bufPool *sync.Pool
	}
	tests := []struct {
		name string
		args args
		want *bytes.Buffer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getBuffer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
