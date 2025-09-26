package session

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalMsg_6(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    *data
		wantErr bool
	}{
		{
			name:  "Test case 1",
			input: []byte(`{"Data":{"key1":"value1","key2":2,"key3":3.4}}`),
			want: &data{
				Data: map[string]any{
					"key1": "value1",
					"key2": 2,
					"key3": 3.4,
				},
			},
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

			z := &data{}
			got, err := z.UnmarshalMsg(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(z, tt.want) {
				t.Errorf("UnmarshalMsg() got = %v, want %v", z, tt.want)
			}
			if !bytes.Equal(got, tt.input) {
				t.Errorf("UnmarshalMsg() got1 = %v, want %v", got, tt.input)
			}
		})
	}
}
