package publisher

import (
	"fmt"
	"testing"
)

func TestWrite_6(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    int
		wantErr bool
	}{
		{
			name:    "Nil input",
			input:   nil,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Empty input",
			input:   []byte{},
			want:    0,
			wantErr: false,
		},
		{
			name:    "Non-empty input",
			input:   []byte("<html><body></body></html>"),
			want:    22,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &htmlElementsCollectorWriter{}
			got, err := w.Write(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("htmlElementsCollectorWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("htmlElementsCollectorWriter.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}
