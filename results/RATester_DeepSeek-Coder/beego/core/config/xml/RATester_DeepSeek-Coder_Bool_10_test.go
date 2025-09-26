package xml

import (
	"fmt"
	"testing"
)

func TestBool_10(t *testing.T) {
	t.Parallel()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "true",
			"key2": "false",
			"key3": "invalid",
		},
	}

	tests := []struct {
		name    string
		key     string
		want    bool
		wantErr bool
	}{
		{"exist true", "key1", true, false},
		{"exist false", "key2", false, false},
		{"not exist", "key3", false, true},
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
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
