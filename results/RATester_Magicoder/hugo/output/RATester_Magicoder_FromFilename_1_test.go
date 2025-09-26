package output

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromFilename_1(t *testing.T) {
	formats := Formats{
		{Name: "html", Path: "html"},
		{Name: "rss", Path: "rss"},
	}

	tests := []struct {
		name     string
		filename string
		want     Format
		found    bool
	}{
		{
			name:     "html",
			filename: "mytemplate.html",
			want:     Format{Name: "html", Path: "html"},
			found:    true,
		},
		{
			name:     "rss",
			filename: "mytemplate.rss",
			want:     Format{Name: "rss", Path: "rss"},
			found:    true,
		},
		{
			name:     "amp",
			filename: "mytemplate.amp.html",
			want:     Format{Name: "html", Path: "html"},
			found:    true,
		},
		{
			name:     "no extension",
			filename: "mytemplate",
			want:     Format{Name: "html", Path: "html"},
			found:    true,
		},
		{
			name:     "not found",
			filename: "mytemplate.unknown",
			want:     Format{},
			found:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, found := formats.FromFilename(tt.filename)
			if !reflect.DeepEqual(got, tt.want) || found != tt.found {
				t.Errorf("FromFilename() = %v, %v, want %v, %v", got, found, tt.want, tt.found)
			}
		})
	}
}
