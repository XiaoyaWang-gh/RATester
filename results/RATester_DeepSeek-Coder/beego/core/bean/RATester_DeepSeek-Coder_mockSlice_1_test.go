package bean

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMockSlice_1(t *testing.T) {
	type testStruct struct {
		Name string
	}
	tests := []struct {
		name     string
		tagValue string
		wantErr  bool
	}{
		{
			name:     "valid tag",
			tagValue: "length:3",
			wantErr:  false,
		},
		{
			name:     "invalid tag",
			tagValue: "length:a",
			wantErr:  true,
		},
		{
			name:     "empty tag",
			tagValue: "",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pvv := reflect.ValueOf(&[]testStruct{})
			err := mockSlice(tt.tagValue, pvv)
			if (err != nil) != tt.wantErr {
				t.Errorf("mockSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
