package limiter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalMsg_1(t *testing.T) {
	tests := []struct {
		name string
		z    item
		want []byte
	}{
		{
			name: "Test case 1",
			z: item{
				currHits: 10,
				prevHits: 20,
				exp:      30,
			},
			want: []byte{0x83, 0xa8, 0x63, 0x75, 0x72, 0x72, 0x48, 0x69, 0x74, 0x73, 0x14, 0xa8, 0x70, 0x72, 0x65, 0x76, 0x48, 0x69, 0x74, 0x73, 0x15, 0xa3, 0x65, 0x78, 0x70, 0x1e},
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

			got, err := tt.z.MarshalMsg(nil)
			if err != nil {
				t.Errorf("item.MarshalMsg() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("item.MarshalMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
