package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseNumberRangePair_1(t *testing.T) {
	type args struct {
		firstRangeStr  string
		secondRangeStr string
	}
	tests := []struct {
		name    string
		args    args
		want    []NumberPair
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				firstRangeStr:  "1-5",
				secondRangeStr: "2-6",
			},
			want: []NumberPair{
				{First: 1, Second: 2},
				{First: 2, Second: 3},
				{First: 3, Second: 4},
				{First: 4, Second: 5},
				{First: 5, Second: 6},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				firstRangeStr:  "1-5",
				secondRangeStr: "2-6-7",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				firstRangeStr:  "1-5-6",
				secondRangeStr: "2-6",
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

			got, err := parseNumberRangePair(tt.args.firstRangeStr, tt.args.secondRangeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseNumberRangePair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNumberRangePair() = %v, want %v", got, tt.want)
			}
		})
	}
}
