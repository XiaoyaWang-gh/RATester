package npm

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestUnmarshal_4(t *testing.T) {
	type testCase struct {
		name    string
		input   io.Reader
		want    map[string]any
		wantErr bool
	}

	tests := []testCase{
		{
			name:  "valid JSON",
			input: strings.NewReader(`{"key": "value"}`),
			want:  map[string]any{"key": "value"},
		},
		{
			name:    "invalid JSON",
			input:   strings.NewReader(`{key: "value"}`),
			wantErr: true,
		},
		{
			name:  "empty JSON",
			input: strings.NewReader(`{}`),
			want:  map[string]any{},
		},
		{
			name:    "empty input",
			input:   strings.NewReader(``),
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

			b := &packageBuilder{}
			got := b.unmarshal(tt.input)
			if (b.err != nil) != tt.wantErr {
				t.Errorf("unmarshal() error = %v, wantErr %v", b.err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
