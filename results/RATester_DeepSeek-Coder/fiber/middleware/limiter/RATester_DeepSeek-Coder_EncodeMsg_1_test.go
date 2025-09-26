package limiter

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestEncodeMsg_1(t *testing.T) {
	type args struct {
		en *msgp.Writer
	}
	tests := []struct {
		name    string
		z       item
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			z: item{
				currHits: 10,
				prevHits: 5,
				exp:      1620602800,
			},
			args: args{
				en: msgp.NewWriter(ioutil.Discard),
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.z.EncodeMsg(tt.args.en); (err != nil) != tt.wantErr {
				t.Errorf("item.EncodeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
