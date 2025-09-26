package tplimpl

import (
	"fmt"
	"testing"
)

func TestErrWithFileContext_2(t *testing.T) {
	type args struct {
		what string
		err  error
	}
	tests := []struct {
		name    string
		info    templateInfo
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

			err := tt.info.errWithFileContext(tt.args.what, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("errWithFileContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
