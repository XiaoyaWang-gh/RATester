package acme

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCheckFile_1(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	tests := []struct {
		name    string
		file    string
		want    bool
		wantErr bool
	}{
		{
			name:    "File exists and is not empty",
			file:    tempFile.Name(),
			want:    true,
			wantErr: false,
		},
		{
			name:    "File does not exist",
			file:    "nonexistentfile",
			want:    false,
			wantErr: true,
		},
		{
			name:    "File exists but is empty",
			file:    tempFile.Name(),
			want:    false,
			wantErr: false,
		},
		{
			name:    "File exists but permissions are too open",
			file:    tempFile.Name(),
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := CheckFile(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
