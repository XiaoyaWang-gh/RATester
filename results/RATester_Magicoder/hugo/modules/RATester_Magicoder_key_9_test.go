package modules

import (
	"fmt"
	"testing"
)

func Testkey_9(t *testing.T) {
	type fields struct {
		Source       string
		Target       string
		Lang         string
		IncludeFiles any
		ExcludeFiles any
		DisableWatch bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				Source: "scss",
				Target: "assets/bootstrap/scss",
				Lang:   "scss",
			},
			want: "scss/scss/assets/bootstrap/scss",
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := Mount{
				Source:       tt.fields.Source,
				Target:       tt.fields.Target,
				Lang:         tt.fields.Lang,
				IncludeFiles: tt.fields.IncludeFiles,
				ExcludeFiles: tt.fields.ExcludeFiles,
				DisableWatch: tt.fields.DisableWatch,
			}
			if got := m.key(); got != tt.want {
				t.Errorf("key() = %v, want %v", got, tt.want)
			}
		})
	}
}
