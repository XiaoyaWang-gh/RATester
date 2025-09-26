package resource

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeByLanguageInterface_1(t *testing.T) {
	type testCase struct {
		name    string
		input   any
		want    any
		wantErr bool
	}

	tests := []testCase{
		{
			name:  "Merge two Resources",
			input: Resources{},
			want:  Resources{},
		},
		{
			name:    "Error when input is not a Resources",
			input:   "not a Resources",
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

			r := Resources{}
			got, err := r.MergeByLanguageInterface(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MergeByLanguageInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeByLanguageInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
