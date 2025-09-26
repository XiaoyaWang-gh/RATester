package main

import (
	"fmt"
	"testing"
)

func TestWriteFile_1(t *testing.T) {
	type args struct {
		name string
		desc *File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				name: "test.go",
				desc: &File{
					Package: "test",
					Imports: []string{"fmt"},
					Elements: []Element{
						{
							Name:  "test",
							Value: "test",
						},
					},
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

			f := fileWriter{
				baseDir: "testdata",
			}
			if err := f.writeFile(tt.args.name, tt.args.desc); (err != nil) != tt.wantErr {
				t.Errorf("writeFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
