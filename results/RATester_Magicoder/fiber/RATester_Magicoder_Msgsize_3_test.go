package fiber

import (
	"fmt"
	"testing"
)

func TestMsgsize_3(t *testing.T) {
	tests := []struct {
		name string
		z    redirectionMsgs
		want int
	}{
		{
			name: "Test Case 1",
			z:    redirectionMsgs{redirectionMsg{key: "key1", value: "value1", level: 1, isOldInput: true}},
			want: 100, // replace with the expected size
		},
		{
			name: "Test Case 2",
			z:    redirectionMsgs{redirectionMsg{key: "key2", value: "value2", level: 2, isOldInput: false}},
			want: 100, // replace with the expected size
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

			if got := tt.z.Msgsize(); got != tt.want {
				t.Errorf("Msgsize() = %v, want %v", got, tt.want)
			}
		})
	}
}
