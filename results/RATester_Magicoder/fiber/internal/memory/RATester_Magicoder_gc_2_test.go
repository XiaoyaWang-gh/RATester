package memory

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Testgc_2(t *testing.T) {
	type fields struct {
		data map[string]item
		sync.RWMutex
	}
	type args struct {
		sleep time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			s := &Storage{
				data: tt.fields.data,
			}
			s.gc(tt.args.sleep)
		})
	}
}
