package paths

import (
	"fmt"
	"testing"
)

func TestFilename_2(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name     string
		args     args
		wantName string
	}{
		{
			name: "Test case 1",
			args: args{
				in: "test.txt",
			},
			wantName: "test",
		},
		{
			name: "Test case 2",
			args: args{
				in: "test.go",
			},
			wantName: "test",
		},
		{
			name: "Test case 3",
			args: args{
				in: "test",
			},
			wantName: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotName := Filename(tt.args.in); gotName != tt.wantName {
				t.Errorf("Filename() = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
