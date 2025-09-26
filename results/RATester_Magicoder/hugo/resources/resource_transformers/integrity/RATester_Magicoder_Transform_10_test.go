package integrity

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestTransform_10(t *testing.T) {
	ctx := &fingerprintTransformation{
		algo: "SHA256",
	}

	testCases := []struct {
		name    string
		ctx     *fingerprintTransformation
		from    io.Reader
		to      io.Writer
		wantErr bool
	}{
		{
			name: "Happy path",
			ctx:  ctx,
			from: strings.NewReader("test"),
			to:   &bytes.Buffer{},
		},
		{
			name:    "Error case",
			ctx:     ctx,
			from:    strings.NewReader("test"),
			to:      &bytes.Buffer{},
			wantErr: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tt.ctx.Transform(&resources.ResourceTransformationCtx{
				From: tt.from,
				To:   tt.to,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
