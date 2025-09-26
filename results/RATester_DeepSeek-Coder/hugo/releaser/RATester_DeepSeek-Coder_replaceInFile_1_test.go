package releaser

import (
	"fmt"
	"testing"
)

func TestReplaceInFile_1(t *testing.T) {
	type args struct {
		filename string
		oldNew   []string
	}
	tests := []struct {
		name    string
		r       *ReleaseHandler
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

			if err := tt.r.replaceInFile(tt.args.filename, tt.args.oldNew...); (err != nil) != tt.wantErr {
				t.Errorf("ReleaseHandler.replaceInFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
