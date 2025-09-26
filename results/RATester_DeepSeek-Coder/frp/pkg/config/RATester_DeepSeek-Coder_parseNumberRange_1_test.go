package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseNumberRange_1(t *testing.T) {
	type args struct {
		firstRangeStr string
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				firstRangeStr: "1-5",
			},
			want:    []int64{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				firstRangeStr: "1-1",
			},
			want:    []int64{1},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				firstRangeStr: "1-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Case 4",
			args: args{
				firstRangeStr: "a-5",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Case 5",
			args: args{
				firstRangeStr: "1-a",
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

			got, err := parseNumberRange(tt.args.firstRangeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseNumberRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNumberRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
