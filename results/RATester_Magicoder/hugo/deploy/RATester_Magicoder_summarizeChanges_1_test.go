package deploy

import (
	"fmt"
	"testing"
)

func TestsummarizeChanges_1(t *testing.T) {
	type args struct {
		uploads []*fileToUpload
		deletes []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				uploads: []*fileToUpload{
					{
						Local: &localFile{
							UploadSize: 1024,
						},
					},
				},
				deletes: []string{"file1", "file2"},
			},
			want: "Identified 1 file(s) to upload, totaling 1.0 kB, and 2 file(s) to delete.",
		},
		{
			name: "Test case 2",
			args: args{
				uploads: []*fileToUpload{
					{
						Local: &localFile{
							UploadSize: 2048,
						},
					},
					{
						Local: &localFile{
							UploadSize: 512,
						},
					},
				},
				deletes: []string{"file3", "file4", "file5"},
			},
			want: "Identified 2 file(s) to upload, totaling 2.5 kB, and 3 file(s) to delete.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := summarizeChanges(tt.args.uploads, tt.args.deletes); got != tt.want {
				t.Errorf("summarizeChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
