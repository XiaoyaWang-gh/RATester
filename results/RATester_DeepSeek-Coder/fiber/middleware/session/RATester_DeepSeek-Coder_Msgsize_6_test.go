package session

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestMsgsize_6(t *testing.T) {
	testCases := []struct {
		name string
		data *data
		want int
	}{
		{
			name: "empty data",
			data: &data{},
			want: 1 + 5 + msgp.MapHeaderSize,
		},
		{
			name: "non-empty data",
			data: &data{
				Data: map[string]any{
					"key1": "value1",
					"key2": 2,
				},
			},
			want: 1 + 5 + msgp.MapHeaderSize +
				msgp.StringPrefixSize + len("key1") + msgp.GuessSize("value1") +
				msgp.StringPrefixSize + len("key2") + msgp.GuessSize(2),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.data.Msgsize()
			if got != tc.want {
				t.Errorf("Msgsize() = %v, want %v", got, tc.want)
			}
		})
	}
}
