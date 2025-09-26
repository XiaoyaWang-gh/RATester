package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/source"
)

func TestGitInfo_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	expected := source.GitInfo{}
	if got := p.GitInfo(); !reflect.DeepEqual(got, expected) {
		t.Errorf("GitInfo() = %v, want %v", got, expected)
	}
}
