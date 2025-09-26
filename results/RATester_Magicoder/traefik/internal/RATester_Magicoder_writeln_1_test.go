package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestWriteln_1(t *testing.T) {
	type args struct {
		a []interface{}
	}
	tests := []struct {
		name string
		ew   *errWriter
		args args
		want error
	}{
		{
			name: "Test case 1",
			ew:   &errWriter{w: &bytes.Buffer{}, err: nil},
			args: args{a: []interface{}{"test"}},
			want: nil,
		},
		{
			name: "Test case 2",
			ew:   &errWriter{w: &bytes.Buffer{}, err: errors.New("test error")},
			args: args{a: []interface{}{"test"}},
			want: errors.New("test error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.ew.writeln(tt.args.a...)
			if !reflect.DeepEqual(tt.ew.err, tt.want) {
				t.Errorf("writeln() error = %v, wantErr %v", tt.ew.err, tt.want)
			}
		})
	}
}
