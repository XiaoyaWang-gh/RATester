package binding

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"testing"
)

func TestsetArrayOfMultipartFormFiles_1(t *testing.T) {
	type testStruct struct {
		Files []*multipart.FileHeader
	}

	testCases := []struct {
		name    string
		input   testStruct
		want    bool
		wantErr error
	}{
		{
			name: "Test case 1",
			input: testStruct{
				Files: []*multipart.FileHeader{
					{
						Filename: "file1.txt",
					},
					{
						Filename: "file2.txt",
					},
				},
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "Test case 2",
			input: testStruct{
				Files: []*multipart.FileHeader{
					{
						Filename: "file1.txt",
					},
				},
			},
			want:    false,
			wantErr: ErrMultiFileHeaderLenInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			value := reflect.ValueOf(tc.input)
			field := reflect.StructField{
				Name: "Files",
				Type: reflect.TypeOf([]*multipart.FileHeader{}),
			}

			got, err := setArrayOfMultipartFormFiles(value, field, tc.input.Files)
			if err != tc.wantErr {
				t.Errorf("setArrayOfMultipartFormFiles() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("setArrayOfMultipartFormFiles() = %v, want %v", got, tc.want)
			}
		})
	}
}
