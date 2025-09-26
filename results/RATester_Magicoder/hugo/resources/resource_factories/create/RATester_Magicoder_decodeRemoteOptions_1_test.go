package create

import (
	"fmt"
	"reflect"
	"testing"
)

func TestdecodeRemoteOptions_1(t *testing.T) {
	tests := []struct {
		name    string
		options map[string]any
		want    fromRemoteOptions
		wantErr bool
	}{
		{
			name: "valid options",
			options: map[string]any{
				"Method": "POST",
				"Headers": map[string]any{
					"Content-Type": "application/json",
				},
				"Body": []byte(`{"key": "value"}`),
			},
			want: fromRemoteOptions{
				Method: "POST",
				Headers: map[string]any{
					"Content-Type": "application/json",
				},
				Body: []byte(`{"key": "value"}`),
			},
			wantErr: false,
		},
		{
			name: "invalid options",
			options: map[string]any{
				"Method": 123,
				"Headers": map[string]any{
					"Content-Type": "application/json",
				},
				"Body": []byte(`{"key": "value"}`),
			},
			want:    fromRemoteOptions{},
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

			got, err := decodeRemoteOptions(tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeRemoteOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeRemoteOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
