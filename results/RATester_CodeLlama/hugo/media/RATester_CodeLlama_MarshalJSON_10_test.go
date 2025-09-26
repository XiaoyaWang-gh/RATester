package media

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalJSON_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Type{
		Type:      "application/rss+xml",
		MainType:  "application",
		SubType:   "rss",
		Delimiter: ".",
		FirstSuffix: SuffixInfo{
			Suffix:     "xml",
			FullSuffix: ".xml",
		},
		mimeSuffix:  "xml",
		SuffixesCSV: "jpg,jpeg",
	}

	b, err := json.Marshal(&m)
	if err != nil {
		t.Fatal(err)
	}

	var got Type
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(m, got) {
		t.Errorf("got %v, want %v", got, m)
	}
}
