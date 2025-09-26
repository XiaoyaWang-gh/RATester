package binding

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSetWithProperType_1(t *testing.T) {
	type testStruct struct {
		IntField    int
		StringField string
		BoolField   bool
		TimeField   time.Time
	}

	testCases := []struct {
		name    string
		val     string
		field   reflect.StructField
		wantErr bool
	}{
		{
			name:    "Test int field",
			val:     "10",
			field:   reflect.TypeOf(testStruct{}).Field(0),
			wantErr: false,
		},
		{
			name:    "Test string field",
			val:     "test",
			field:   reflect.TypeOf(testStruct{}).Field(1),
			wantErr: false,
		},
		{
			name:    "Test bool field",
			val:     "true",
			field:   reflect.TypeOf(testStruct{}).Field(2),
			wantErr: false,
		},
		{
			name:    "Test time field",
			val:     "2022-01-01T00:00:00Z",
			field:   reflect.TypeOf(testStruct{}).Field(3),
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			value := reflect.ValueOf(&testStruct{}).Elem().FieldByName(tc.field.Name)
			err := setWithProperType(tc.val, value, tc.field)
			if (err != nil) != tc.wantErr {
				t.Errorf("setWithProperType() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
