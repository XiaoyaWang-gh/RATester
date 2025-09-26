package csrf

import (
	"fmt"
	"sync"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestEncodeMsg_5(t *testing.T) {
	type testCase struct {
		name    string
		manager storageManager
		wantErr bool
	}

	tests := []testCase{
		{
			name: "Test case 1",
			manager: storageManager{
				pool: sync.Pool{},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			manager: storageManager{
				pool: sync.Pool{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &msgp.Writer{}
			err := tt.manager.EncodeMsg(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("storageManager.EncodeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
