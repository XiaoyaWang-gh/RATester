package highlight

import (
	"fmt"
	"testing"
)

func TestGetPreWrapper_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	language := "en"
	writeCounter := &byteCountFlexiWriter{}
	preWrapper := getPreWrapper(language, writeCounter)
	if preWrapper.language != language {
		t.Errorf("language is not equal")
	}
	if preWrapper.writeCounter != writeCounter {
		t.Errorf("writeCounter is not equal")
	}
}
