package allconfig

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/herrors"
)

func TestWrapFileError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := configLoader{}
	err := errors.New("test error")
	filename := "test.yaml"
	want := herrors.NewFileErrorFromFile(err, filename, l.Fs, nil)
	got := l.wrapFileError(err, filename)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
