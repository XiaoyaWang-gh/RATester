package dynacache

import (
	"fmt"
	"testing"
	"time"
)

func TestStart_1(t *testing.T) {
	t.Parallel()

	type fields struct {
		opts Options
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test start method",
			fields: fields{
				opts: Options{
					CheckInterval: 1 * time.Second,
					MaxSize:       100,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Cache{
				opts: tt.fields.opts,
			}

			stop := c.start()
			time.Sleep(2 * time.Second)
			stop()
		})
	}
}
