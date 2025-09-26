package filesystems

import (
	"fmt"
	"testing"
)

func TestMakePathRelative_3(t *testing.T) {
	type args struct {
		filename    string
		checkExists bool
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Test case 1",
			args: args{
				filename:    "test.txt",
				checkExists: true,
			},
			want:  "test.txt",
			want1: true,
		},
		{
			name: "Test case 2",
			args: args{
				filename:    "nonexistent.txt",
				checkExists: true,
			},
			want:  "",
			want1: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &SourceFilesystem{}
			got, got1 := d.MakePathRelative(tt.args.filename, tt.args.checkExists)
			if got != tt.want {
				t.Errorf("MakePathRelative() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakePathRelative() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
