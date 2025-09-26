package integrity

import (
	"fmt"
	"testing"
)

func Testintegrity_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testCases := []struct {
		algo string
		sum  []byte
		want string
	}{
		{
			algo: "SHA256",
			sum:  []byte("test"),
			want: "SHA256-dGVzdA==",
		},
		{
			algo: "SHA1",
			sum:  []byte("hello"),
			want: "SHA1-aGVsbG8=",
		},
		{
			algo: "MD5",
			sum:  []byte("world"),
			want: "MD5-d29ybGQ=",
		},
	}

	for _, tc := range testCases {
		got := integrity(tc.algo, tc.sum)
		if got != tc.want {
			t.Errorf("integrity(%s, %v) = %s, want %s", tc.algo, tc.sum, got, tc.want)
		}
	}
}
