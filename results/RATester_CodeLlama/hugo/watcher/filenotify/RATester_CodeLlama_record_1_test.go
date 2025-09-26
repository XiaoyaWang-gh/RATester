package filenotify

import (
	"fmt"
	"testing"
)

func TestRecord_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &recording{}
	filename := "testdata/test.txt"
	err := r.record(filename)
	if err != nil {
		t.Fatal(err)
	}

	if r.FileInfo == nil {
		t.Fatal("FileInfo is nil")
	}

	if r.FileInfo.Name() != "test.txt" {
		t.Fatalf("FileInfo.Name() = %s, want %s", r.FileInfo.Name(), "test.txt")
	}

	if r.FileInfo.Size() != 10 {
		t.Fatalf("FileInfo.Size() = %d, want %d", r.FileInfo.Size(), 10)
	}

	if r.FileInfo.Mode() != 0 {
		t.Fatalf("FileInfo.Mode() = %d, want %d", r.FileInfo.Mode(), 0)
	}

	if r.FileInfo.ModTime().Unix() != 1619781200 {
		t.Fatalf("FileInfo.ModTime() = %d, want %d", r.FileInfo.ModTime().Unix(), 1619781200)
	}

	if r.FileInfo.IsDir() {
		t.Fatal("FileInfo.IsDir() = true, want false")
	}

	if r.entries == nil {
		t.Fatal("entries is nil")
	}

	if len(r.entries) != 0 {
		t.Fatalf("len(entries) = %d, want %d", len(r.entries), 0)
	}
}
