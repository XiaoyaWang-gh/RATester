package hugolib

import (
	"fmt"
	"testing"
)

func TestcreateShortcodePlaceholder_1(t *testing.T) {
	type args struct {
		sid     string
		id      uint64
		ordinal int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				sid:     "SID",
				id:      1,
				ordinal: 1,
			},
			want: "HBHB1SID1HBHB",
		},
		{
			name: "Test Case 2",
			args: args{
				sid:     "SID",
				id:      2,
				ordinal: 2,
			},
			want: "HBHB2SID2HBHB",
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

			if got := createShortcodePlaceholder(tt.args.sid, tt.args.id, tt.args.ordinal); got != tt.want {
				t.Errorf("createShortcodePlaceholder() = %v, want %v", got, tt.want)
			}
		})
	}
}
