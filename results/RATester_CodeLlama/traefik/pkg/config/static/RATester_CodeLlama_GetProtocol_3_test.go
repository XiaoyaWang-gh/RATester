package static

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetProtocol_3(t *testing.T) {
	testCases := []struct {
		desc     string
		ep       EntryPoint
		expected string
		err      error
	}{
		{
			desc: "should return tcp",
			ep: EntryPoint{
				Address: "127.0.0.1:80",
			},
			expected: "tcp",
		},
		{
			desc: "should return tcp",
			ep: EntryPoint{
				Address: "127.0.0.1:80/tcp",
			},
			expected: "tcp",
		},
		{
			desc: "should return udp",
			ep: EntryPoint{
				Address: "127.0.0.1:80/udp",
			},
			expected: "udp",
		},
		{
			desc: "should return error",
			ep: EntryPoint{
				Address: "127.0.0.1:80/foo",
			},
			err: fmt.Errorf("invalid protocol: foo"),
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			actual, err := test.ep.GetProtocol()
			if test.err != nil {
				require.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, actual)
			}
		})
	}
}
