package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestBoolValue_1(t *testing.T) {
	tests := []struct {
		name           string
		sourceParam    string
		valueMustExist bool
		valueFunc      func(sourceParam string) string
		errorFunc      func(sourceParam string, values []string, message interface{}, internalError error) error
		wantErr        bool
	}{
		{
			name:           "Test case 1",
			sourceParam:    "field name",
			valueMustExist: true,
			valueFunc: func(sourceParam string) string {
				if sourceParam == "field name" {
					return "true"
				}
				return ""
			},
			errorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
				return errors.New("error")
			},
			wantErr: false,
		},
		{
			name:           "Test case 2",
			sourceParam:    "field name",
			valueMustExist: true,
			valueFunc: func(sourceParam string) string {
				if sourceParam == "field name" {
					return ""
				}
				return ""
			},
			errorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
				return errors.New("error")
			},
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

			b := &ValueBinder{
				ValueFunc: tt.valueFunc,
				ErrorFunc: tt.errorFunc,
			}
			if err := b.boolValue(tt.sourceParam, nil, tt.valueMustExist); (err != nil) != tt.wantErr {
				t.Errorf("boolValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
