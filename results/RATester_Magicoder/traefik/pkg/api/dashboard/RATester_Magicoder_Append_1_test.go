package dashboard

import (
	"fmt"
	"io/fs"
	"testing"

	"github.com/gorilla/mux"
)

func TestAppend_1(t *testing.T) {
	type args struct {
		router       *mux.Router
		customAssets fs.FS
	}
	tests := []struct {
		name string
		args args
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

			Append(tt.args.router, tt.args.customAssets)
		})
	}
}
