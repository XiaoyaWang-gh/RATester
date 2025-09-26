package session

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestVisit_1(t *testing.T) {
	type args struct {
		paths string
		f     fs.FileInfo
		err   error
	}
	tests := []struct {
		name    string
		as      *activeSession
		args    args
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

			as := &activeSession{
				total: tt.as.total,
			}
			if err := as.visit(tt.args.paths, tt.args.f, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("visit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
