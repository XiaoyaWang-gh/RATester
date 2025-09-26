package fiber

import (
	"fmt"
	"testing"
)

func TestMarshalMsg_3(t *testing.T) {
	tests := []struct {
		name    string
		z       redirectionMsgs
		wantErr bool
	}{
		{
			name: "Test case 1",
			z: []redirectionMsg{
				{
					key:        "key1",
					value:      "value1",
					level:      1,
					isOldInput: true,
				},
				{
					key:        "key2",
					value:      "value2",
					level:      2,
					isOldInput: false,
				},
			},
			wantErr: false,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := tt.z.MarshalMsg(nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("redirectionMsgs.MarshalMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
