package crypt

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPKCS5Padding_1(t *testing.T) {
	t.Run("TestPKCS5Padding", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name       string
			ciphertext []byte
			blockSize  int
			want       []byte
		}{
			{
				name:       "Test Case 1",
				ciphertext: []byte("hello"),
				blockSize:  8,
				want:       []byte("hello\x03\x03\x03"),
			},
			{
				name:       "Test Case 2",
				ciphertext: []byte("hello"),
				blockSize:  5,
				want:       []byte("hello"),
			},
			{
				name:       "Test Case 3",
				ciphertext: []byte("hello"),
				blockSize:  1,
				want:       []byte("hello\x05"),
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				got := PKCS5Padding(tc.ciphertext, tc.blockSize)
				if !bytes.Equal(got, tc.want) {
					t.Errorf("Expected %v, got %v", tc.want, got)
				}
			})
		}
	})
}
