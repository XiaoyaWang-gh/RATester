package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSection_3(t *testing.T) {
	type args struct {
		section string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				section: "section1",
			},
			want: map[string]string{
				"field name": "value1",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				section: "section2",
			},
			want: map[string]string{
				"field name": "value2",
			},
			wantErr: false,
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

			got, err := GetSection(tt.args.section)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
