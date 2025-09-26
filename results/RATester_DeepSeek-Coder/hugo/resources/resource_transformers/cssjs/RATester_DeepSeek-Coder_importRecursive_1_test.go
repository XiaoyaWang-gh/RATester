package cssjs

import (
	"fmt"
	"testing"
)

func TestImportRecursive_1(t *testing.T) {
	type args struct {
		lineNum int
		content string
		inPath  string
	}
	tests := []struct {
		name    string
		imp     *importResolver
		args    args
		want    int
		want1   string
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

			got, got1, err := tt.imp.importRecursive(tt.args.lineNum, tt.args.content, tt.args.inPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("importRecursive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("importRecursive() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("importRecursive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
