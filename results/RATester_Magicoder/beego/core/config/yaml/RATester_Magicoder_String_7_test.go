package yaml

import (
	"fmt"
	"testing"
)

func TestString_7(t *testing.T) {
	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key": "value",
		},
	}

	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "test case 1",
			key:     "key",
			want:    "value",
			wantErr: false,
		},
		{
			name:    "test case 2",
			key:     "not_exist",
			want:    "",
			wantErr: false,
		},
		{
			name:    "test case 3",
			key:     "key",
			want:    "value",
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

			got, err := cc.String(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
