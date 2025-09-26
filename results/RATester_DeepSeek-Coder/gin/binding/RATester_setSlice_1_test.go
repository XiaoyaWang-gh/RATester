package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetSlice_1(t *testing.T) {
	type TestStruct struct {
		Values []string
	}

	testCases := []struct {
		name    string
		vals    []string
		want    TestStruct
		wantErr bool
	}{
		{
			name: "Test case 1",
			vals: []string{"val1", "val2", "val3"},
			want: TestStruct{
				Values: []string{"val1", "val2", "val3"},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			vals: []string{"val4", "val5", "val6"},
			want: TestStruct{
				Values: []string{"val4", "val5", "val6"},
			},
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

			ts := TestStruct{}
			value := reflect.ValueOf(&ts).Elem().FieldByName("Values")
			field, _ := reflect.TypeOf(ts).FieldByName("Values")
			err := setSlice(tt.vals, value, field)
			if (err != nil) != tt.wantErr {
				t.Errorf("setSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(ts, tt.want) {
				t.Errorf("setSlice() = %v, want %v", ts, tt.want)
			}
		})
	}
}
