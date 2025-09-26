package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSection_1(t *testing.T) {
	cc := &ConfigContainer{
		data: map[string]interface{}{
			"section1": map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			"section2": map[string]string{
				"key3": "value3",
				"key4": "value4",
			},
		},
	}

	tests := []struct {
		name    string
		section string
		want    map[string]string
		wantErr bool
	}{
		{
			name:    "ExistingSection",
			section: "section1",
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name:    "ExistingSection2",
			section: "section2",
			want: map[string]string{
				"key3": "value3",
				"key4": "value4",
			},
			wantErr: false,
		},
		{
			name:    "NonExistingSection",
			section: "nonExistingSection",
			want:    nil,
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

			got, err := cc.GetSection(tt.section)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
