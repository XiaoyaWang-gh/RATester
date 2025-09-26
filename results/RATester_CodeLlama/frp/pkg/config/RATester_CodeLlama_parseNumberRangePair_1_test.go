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
			name: "test case 1",
			args: args{
				firstRangeStr:  "1-2",
				secondRangeStr: "3-4",
			},
			want: []NumberPair{
				{
					First:  1,
					Second: 3,
				},
				{
					First:  2,
					Second: 4,
				},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				firstRangeStr:  "1-2",
				secondRangeStr: "3-4-5",
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
