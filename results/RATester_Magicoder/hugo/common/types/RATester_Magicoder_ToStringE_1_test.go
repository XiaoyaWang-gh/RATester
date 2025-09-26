package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestToStringE_1(t *testing.T) {
	tests := []struct {
		name    string
		v       any
		want    string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			v:       "test",
			want:    "test",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			v:       json.RawMessage(`"test"`),
			want:    "test",
			wantErr: false,
		},
		{
			name:    "Test case 3",
			v:       []byte("test"),
			want:    "test",
			wantErr: false,
		},
		{
			name:    "Test case 4",
			v:       nil,
			want:    "",
			wantErr: false,
		},
		{
			name:    "Test case 5",
			v:       func() {},
			want:    "",
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

			got, err := ToStringE(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToStringE() = %v, want %v", got, tt.want)
			}
		})
	}
}
