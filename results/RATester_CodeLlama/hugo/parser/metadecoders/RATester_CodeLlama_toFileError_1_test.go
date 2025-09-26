package metadecoders

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/herrors"
)

func TestToFileError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := Format("")
	data := []byte("")
	err := errors.New("")
	want := herrors.NewFileErrorFromName(err, "_stream.Format").UpdateContent(bytes.NewReader(data), nil)
	if got := toFileError(f, data, err); !reflect.DeepEqual(got, want) {
		t.Errorf("toFileError() = %v, want %v", got, want)
	}
}
