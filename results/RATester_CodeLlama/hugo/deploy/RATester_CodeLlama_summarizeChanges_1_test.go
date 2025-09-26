package deploy

import (
	"fmt"
	"testing"
)

func TestSummarizeChanges_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	uploads := []*fileToUpload{
		{
			Local: &localFile{
				NativePath: "path/to/file1",
				UploadSize: 100,
			},
		},
		{
			Local: &localFile{
				NativePath: "path/to/file2",
				UploadSize: 200,
			},
		},
	}
	deletes := []string{"path/to/file3"}
	expected := "Identified 2 file(s) to upload, totaling 300B, and 1 file(s) to delete."
	actual := summarizeChanges(uploads, deletes)
	if actual != expected {
		t.Errorf("summarizeChanges(%v, %v) = %q; want %q", uploads, deletes, actual, expected)
	}
}
