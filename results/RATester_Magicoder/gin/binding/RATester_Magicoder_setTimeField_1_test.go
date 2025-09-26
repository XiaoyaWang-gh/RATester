package binding

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestsetTimeField_1(t *testing.T) {
	type testStruct struct {
		TimeField time.Time `time_format:"2006-01-02"`
	}

	testCases := []struct {
		name    string
		val     string
		wantErr bool
		want    time.Time
	}{
		{
			name:    "valid time",
			val:     "2022-01-01",
			wantErr: false,
			want:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:    "invalid time",
			val:     "invalid",
			wantErr: true,
			want:    time.Time{},
		},
		{
			name:    "empty time",
			val:     "",
			wantErr: false,
			want:    time.Time{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := testStruct{}
			v := reflect.ValueOf(&s).Elem().FieldByName("TimeField")
			err := setTimeField(tc.val, reflect.StructField{}, v)
			if (err != nil) != tc.wantErr {
				t.Errorf("setTimeField() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !v.Interface().(time.Time).Equal(tc.want) {
				t.Errorf("setTimeField() = %v, want %v", v.Interface().(time.Time), tc.want)
			}
		})
	}
}
