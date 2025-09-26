package fiber

import (
	"fmt"
	"testing"
)

func TestSendFile_1(t *testing.T) {
	t.Parallel()

	app := New()

	testCases := []struct {
		name    string
		file    string
		config  SendFile
		wantErr bool
	}{
		{
			name:   "Test with valid file",
			file:   "testdata/test.txt",
			config: SendFile{},
		},
		{
			name:    "Test with invalid file",
			file:    "testdata/invalid.txt",
			config:  SendFile{},
			wantErr: true,
		},
		{
			name:    "Test with empty file",
			file:    "",
			config:  SendFile{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := app.AcquireCtx(nil)
			err := ctx.SendFile(tc.file, tc.config)
			if (err != nil) != tc.wantErr {
				t.Errorf("SendFile() error = %v, wantErr %v", err, tc.wantErr)
			}
			app.ReleaseCtx(ctx)
		})
	}
}
