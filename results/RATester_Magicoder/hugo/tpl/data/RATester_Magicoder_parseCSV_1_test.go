package data

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseCSV_1(t *testing.T) {
	type args struct {
		c   []byte
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				c:   []byte("a,b,c\n1,2,3\n4,5,6"),
				sep: ",",
			},
			want:    [][]string{{"a", "b", "c"}, {"1", "2", "3"}, {"4", "5", "6"}},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				c:   []byte("a;b;c\n1;2;3\n4;5;6"),
				sep: ";",
			},
			want:    [][]string{{"a", "b", "c"}, {"1", "2", "3"}, {"4", "5", "6"}},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				c:   []byte("a,b,c\n1,2,3\n4,5,6"),
				sep: " ",
			},
			want:    nil,
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

			got, err := parseCSV(tt.args.c, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
