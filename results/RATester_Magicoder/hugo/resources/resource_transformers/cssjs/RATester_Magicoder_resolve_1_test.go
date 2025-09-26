package cssjs

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func Testresolve_1(t *testing.T) {
	type args struct {
		r      io.Reader
		inPath string
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
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

			imp := &importResolver{
				r:      tt.args.r,
				inPath: tt.args.inPath,
			}
			got, err := imp.resolve()
			if (err != nil) != tt.wantErr {
				t.Errorf("importResolver.resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importResolver.resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
