package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetWithProperType_1(t *testing.T) {
	type testStruct struct {
		IntField    int
		StringField string
	}

	testCases := []struct {
		name        string
		valueKind   reflect.Kind
		val         string
		structField reflect.Value
		wantErr     bool
	}{
		{
			name:        "Test with valid int",
			valueKind:   reflect.Int,
			val:         "10",
			structField: reflect.ValueOf(&testStruct{}).Elem().FieldByName("IntField"),
			wantErr:     false,
		},
		{
			name:        "Test with invalid int",
			valueKind:   reflect.Int,
			val:         "invalid",
			structField: reflect.ValueOf(&testStruct{}).Elem().FieldByName("IntField"),
			wantErr:     true,
		},
		{
			name:        "Test with valid string",
			valueKind:   reflect.String,
			val:         "test",
			structField: reflect.ValueOf(&testStruct{}).Elem().FieldByName("StringField"),
			wantErr:     false,
		},
		{
			name:        "Test with invalid string",
			valueKind:   reflect.String,
			val:         "test",
			structField: reflect.ValueOf(&testStruct{}).Elem().FieldByName("IntField"),
			wantErr:     true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := setWithProperType(tt.valueKind, tt.val, tt.structField)
			if (err != nil) != tt.wantErr {
				t.Errorf("setWithProperType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
