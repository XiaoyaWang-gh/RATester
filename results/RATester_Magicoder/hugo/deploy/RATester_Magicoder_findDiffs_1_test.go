package deploy

import (
	"fmt"
	"testing"

	"gocloud.dev/blob"
)

func TestfindDiffs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Deployer{}

	localFiles := map[string]*localFile{
		"file1": {
			NativePath: "file1",
			UploadSize: 100,
			md5:        []byte{1, 2, 3},
		},
		"file2": {
			NativePath: "file2",
			UploadSize: 200,
			md5:        []byte{4, 5, 6},
		},
	}

	remoteFiles := map[string]*blob.ListObject{
		"file1": {
			Size: 100,
			MD5:  []byte{1, 2, 3},
		},
		"file3": {
			Size: 300,
			MD5:  []byte{7, 8, 9},
		},
	}

	uploads, deletes := d.findDiffs(localFiles, remoteFiles, false)

	if len(uploads) != 1 {
		t.Errorf("Expected 1 upload, got %d", len(uploads))
	}

	if uploads[0].Local.NativePath != "file2" {
		t.Errorf("Expected upload for file2, got %s", uploads[0].Local.NativePath)
	}

	if len(deletes) != 1 {
		t.Errorf("Expected 1 delete, got %d", len(deletes))
	}

	if deletes[0] != "file3" {
		t.Errorf("Expected delete for file3, got %s", deletes[0])
	}
}
