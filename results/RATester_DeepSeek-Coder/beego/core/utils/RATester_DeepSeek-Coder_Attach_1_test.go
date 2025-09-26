package utils

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestAttach_1(t *testing.T) {
	type args struct {
		r        io.Reader
		filename string
		args     []string
	}
	tests := []struct {
		name    string
		e       *Email
		args    args
		wantA   *Attachment
		wantErr bool
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

			gotA, err := tt.e.Attach(tt.args.r, tt.args.filename, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Email.Attach() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("Email.Attach() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}
