package context

import (
	"fmt"
	"testing"
)

func TestParam_2(t *testing.T) {
	input := &BeegoInput{
		pnames:  []string{"key1", "key2", "key3"},
		pvalues: []string{"value1", "value2", "value3"},
	}

	tests := []struct {
		name  string
		input *BeegoInput
		key   string
		want  string
	}{
		{
			name:  "TestParam_1",
			input: input,
			key:   "key1",
			want:  "value1",
		},
		{
			name:  "TestParam_2",
			input: input,
			key:   "key2",
			want:  "value2",
		},
		{
			name:  "TestParam_3",
			input: input,
			key:   "key3",
			want:  "value3",
		},
		{
			name:  "TestParam_4",
			input: input,
			key:   "key4",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.input.Param(tt.key); got != tt.want {
				t.Errorf("Param() = %v, want %v", got, tt.want)
			}
		})
	}
}
