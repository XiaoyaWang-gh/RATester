package yaml

import (
	"fmt"
	"testing"
)

func TestBool_1(t *testing.T) {
	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key": "true",
		},
	}

	tests := []struct {
		name    string
		key     string
		want    bool
		wantErr bool
	}{
		{
			name:    "TestBool_True",
			key:     "key",
			want:    true,
			wantErr: false,
		},
		{
			name:    "TestBool_False",
			key:     "key2",
			want:    false,
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

			got, err := cc.Bool(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bool() got = %v, want %v", got, tt.want)
			}
		})
	}
}
