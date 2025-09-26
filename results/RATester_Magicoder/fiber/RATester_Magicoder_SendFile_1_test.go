package fiber

import (
	"fmt"
	"testing"
	"time"
)

func TestSendFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		app: &App{
			sendfiles: []*sendFileStore{
				{
					config: SendFile{
						CacheDuration: 10 * time.Second,
					},
				},
			},
		},
	}

	err := ctx.SendFile("test.txt")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
