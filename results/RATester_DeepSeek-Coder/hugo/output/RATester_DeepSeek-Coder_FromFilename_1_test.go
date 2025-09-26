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
		{Name: "amp", Path: "amp"},
	}

	tests := []struct {
		name      string
		filename  string
		wantF     Format
		wantFound bool
	}{
		{"html", "mytemplate.html", formats[0], true},
		{"rss", "mytemplate.rss", formats[1], true},
		{"amp", "mytemplate.amp.html", formats[2], true},
		{"no extension", "mytemplate", Format{}, false},
		{"unknown format", "mytemplate.unknown", Format{}, false},
		{"multiple dots", "mytemplate.rss.html", formats[1], true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotF, gotFound := formats.FromFilename(tt.filename)
			if !reflect.DeepEqual(gotF, tt.wantF) || gotFound != tt.wantFound {
				t.Errorf("FromFilename() = %v, %v, want %v, %v", gotF, gotFound, tt.wantF, tt.wantFound)
			}
		})
	}
}
