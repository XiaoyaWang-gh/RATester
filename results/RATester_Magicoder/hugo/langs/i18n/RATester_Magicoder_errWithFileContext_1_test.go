package i18n

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/source"
)

func TesterrWithFileContext_1(t *testing.T) {
	type args struct {
		inerr error
		r     *source.File
	}
	tests := []struct {
		name    string
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

			if err := errWithFileContext(tt.args.inerr, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("errWithFileContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
