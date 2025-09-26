package redactor

import (
	"fmt"
	"testing"
)

func TestAnonymize_1(t *testing.T) {
	type args struct {
		baseConfig interface{}
		indent     bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				baseConfig: map[string]interface{}{
					"field name": "value",
				},
				indent: true,
			},
			want:    "{\n  \"field name\": \"value\"\n}",
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				baseConfig: map[string]interface{}{
					"field name": "value",
				},
				indent: false,
			},
			want:    "{\"field name\":\"value\"}",
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				baseConfig: "invalid",
				indent:     true,
			},
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

			got, err := anonymize(tt.args.baseConfig, tt.args.indent)
			if (err != nil) != tt.wantErr {
				t.Errorf("anonymize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("anonymize() = %v, want %v", got, tt.want)
			}
		})
	}
}
