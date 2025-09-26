package install

import (
	"fmt"
	"testing"
)

func TestCopyFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	src := "./test.txt"
	dest := "./test2.txt"
	w, err := copyFile(src, dest)
	if err != nil {
		t.Errorf("copyFile error:%v", err)
	}
	if w != 10 {
		t.Errorf("copyFile error:%v", w)
	}
}
