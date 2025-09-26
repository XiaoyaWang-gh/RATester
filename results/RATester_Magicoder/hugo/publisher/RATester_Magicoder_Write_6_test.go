package publisher

import (
	"fmt"
	"testing"
)

func TestWrite_6(t *testing.T) {
	tests := []struct {
		name    string
		w       *htmlElementsCollectorWriter
		args    []byte
		want    int
		wantErr bool
	}{
		{
			name: "Test case 1",
			w: &htmlElementsCollectorWriter{
				collector: &htmlElementsCollector{
					elementSet: make(map[string]bool),
				},
			},
			args: []byte("<div>Hello, World!</div>"),
			want: 20,
		},
		{
			name: "Test case 2",
			w: &htmlElementsCollectorWriter{
				collector: &htmlElementsCollector{
					elementSet: make(map[string]bool),
				},
			},
			args: []byte("<div>Hello, World!</div><p>Another paragraph</p>"),
			want: 40,
		},
		{
			name: "Test case 3",
			w: &htmlElementsCollectorWriter{
				collector: &htmlElementsCollector{
					elementSet: make(map[string]bool),
				},
			},
			args: []byte("<div>Hello, World!</div><p>Another paragraph</p><span>A span</span>"),
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.w.Write(tt.args)
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
