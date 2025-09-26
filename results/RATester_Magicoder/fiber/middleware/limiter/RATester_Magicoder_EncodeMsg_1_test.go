package limiter

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestEncodeMsg_1(t *testing.T) {
	tests := []struct {
		name    string
		item    item
		wantErr bool
	}{
		{
			name: "Test case 1",
			item: item{
				currHits: 10,
				prevHits: 5,
				exp:      1640995200,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			item: item{
				currHits: -1,
				prevHits: 0,
				exp:      0,
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			item: item{
				currHits: 0,
				prevHits: 0,
				exp:      0,
			},
			wantErr: false,
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
			if err := tt.item.EncodeMsg(w); (err != nil) != tt.wantErr {
				t.Errorf("item.EncodeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
