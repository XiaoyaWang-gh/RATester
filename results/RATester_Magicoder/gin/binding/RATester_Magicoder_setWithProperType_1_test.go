package binding

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestsetWithProperType_1(t *testing.T) {
	type TestStruct struct {
		Field1 int
		Field2 string
		Field3 time.Time
	}

	testCases := []struct {
		name    string
		val     string
		field   reflect.StructField
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			val:     "10",
			field:   reflect.StructField{Name: "Field1", Type: reflect.TypeOf(10)},
			wantErr: false,
		},
		{
			name:    "Test Case 2",
			val:     "test",
			field:   reflect.StructField{Name: "Field2", Type: reflect.TypeOf("")},
			wantErr: false,
		},
		{
			name:    "Test Case 3",
			val:     "2022-01-01T00:00:00Z",
			field:   reflect.StructField{Name: "Field3", Type: reflect.TypeOf(time.Time{})},
			wantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			testStruct := TestStruct{}
			value := reflect.ValueOf(&testStruct).Elem().FieldByName(tt.field.Name)
			err := setWithProperType(tt.val, value, tt.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("setWithProperType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
