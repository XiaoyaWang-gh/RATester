package webp

import (
	"fmt"
	"image"
	"io"
	"testing"

	"github.com/bep/gowebp/libwebp/webpoptions"
	"github.com/gohugoio/hugo/common/herrors"
)

func TestEncode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var w io.Writer
	var m image.Image
	var o webpoptions.EncodingOptions

	// Act
	err := Encode(w, m, o)

	// Assert
	if err != herrors.ErrFeatureNotAvailable {
		t.Errorf("Encode() = %v, want %v", err, herrors.ErrFeatureNotAvailable)
	}
}
