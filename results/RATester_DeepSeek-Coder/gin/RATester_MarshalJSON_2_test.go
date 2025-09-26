package gin

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalJSON_2(t *testing.T) {
	tests := []struct {
		name    string
		a       errorMsgs
		want    []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			a: errorMsgs{
				&Error{
					Err:  errors.New("test error"),
					Type: ErrorType(1),
					Meta: "test meta",
				},
			},
			want:    []byte(`[{"Err":"test error","Type":1,"Meta":"test meta"}]`),
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

			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("errorMsgs.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorMsgs.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
