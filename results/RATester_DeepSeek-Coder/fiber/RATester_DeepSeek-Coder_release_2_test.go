package fiber

import (
	"fmt"
	"testing"
)

func TestRelease_2(t *testing.T) {
	type fields struct {
		c        *DefaultCtx
		messages redirectionMsgs
		status   int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &Redirect{
				c:        tt.fields.c,
				messages: tt.fields.messages,
				status:   tt.fields.status,
			}
			r.release()
		})
	}
}
