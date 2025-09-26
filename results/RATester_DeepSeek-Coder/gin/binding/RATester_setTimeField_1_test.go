package binding

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSetTimeField_1(t *testing.T) {
	type testStruct struct {
		TimeField time.Time `time_format:"2006-01-02 15:04:05"`
	}

	testCases := []struct {
		name    string
		val     string
		wantErr bool
		want    time.Time
	}{
		{
			name:    "valid time",
			val:     "2022-01-02 15:04:05",
			wantErr: false,
			want:    time.Date(2022, 1, 2, 15, 4, 5, 0, time.Local),
		},
		{
			name:    "empty time",
			val:     "",
			wantErr: false,
			want:    time.Time{},
		},
		{
			name:    "invalid time",
			val:     "2022-13-02 15:04:05",
			wantErr: true,
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
			err := setTimeField(tc.val, v.Type().Field(0), v)
			if (err != nil) != tc.wantErr {
				t.Errorf("setTimeField() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr && s.TimeField != tc.want {
				t.Errorf("setTimeField() = %v, want %v", s.TimeField, tc.want)
			}
		})
	}
}
