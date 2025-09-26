package exif

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/bep/imagemeta"
	"github.com/stretchr/testify/require"
)

func TestDecode_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Decoder{
		includeFieldsRe: regexp.MustCompile(".*"),
	}

	ex, err := d.Decode("test.jpg", imagemeta.JPEG, strings.NewReader(""))
	require.NoError(t, err)
	require.NotNil(t, ex)
}
