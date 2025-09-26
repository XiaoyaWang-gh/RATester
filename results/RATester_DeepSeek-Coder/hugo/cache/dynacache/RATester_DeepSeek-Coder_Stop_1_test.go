package dynacache

import (
	"fmt"
	"sync"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestStop_1(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stopFunc := func() {
		// mock stop function
	}

	tests := []struct {
		name string
		stop func()
	}{
		{
			name: "stop is called once",
			stop: stopFunc,
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
				stopOnce: sync.Once{},
				stop:     tt.stop,
			}

			stopCalled := false
			c.stop = func() {
				stopCalled = true
			}

			c.Stop()
			c.Stop()

			if !stopCalled {
				t.Errorf("Expected stop to be called, but it wasn't")
			}
		})
	}
}
