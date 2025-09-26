package resources

import (
	"fmt"
	"testing"
)

func TestOpenPublishFileForWriting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var r *genericResource
	var relTargetPath string
	// when
	actual, actualErr := r.openPublishFileForWriting(relTargetPath)
	// then
	if actual != nil {
		t.Errorf("actual %v should be nil", actual)
	}
	if actualErr != nil {
		t.Errorf("actual error %v should be nil", actualErr)
	}
}
