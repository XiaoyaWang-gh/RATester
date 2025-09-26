package binding

import (
	"fmt"
	"mime/multipart"
	"net/textproto"
	"reflect"
	"testing"
)

func TestSetArrayOfMultipartFormFiles_1(t *testing.T) {
	type testStruct struct {
		Files []*multipart.FileHeader
	}

	testCases := []struct {
		name     string
		input    testStruct
		expected error
	}{
		{
			name: "Test case 1",
			input: testStruct{
				Files: []*multipart.FileHeader{
					{
						Filename: "file1.txt",
						Size:     1024,
						Header:   make(textproto.MIMEHeader),
					},
					{
						Filename: "file2.txt",
						Size:     2048,
						Header:   make(textproto.MIMEHeader),
					},
				},
			},
			expected: nil,
		},
		{
			name: "Test case 2",
			input: testStruct{
				Files: []*multipart.FileHeader{
					{
						Filename: "file1.txt",
						Size:     1024,
						Header:   make(textproto.MIMEHeader),
					},
				},
			},
			expected: ErrMultiFileHeaderLenInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			value := reflect.ValueOf(&tc.input).Elem().FieldByName("Files")
			field := reflect.StructField{
				Name: "Files",
				Type: reflect.TypeOf([]*multipart.FileHeader{}),
			}
			_, err := setArrayOfMultipartFormFiles(value, field, tc.input.Files)
			if err != tc.expected {
				t.Errorf("Expected error %v, but got %v", tc.expected, err)
			}
		})
	}
}
