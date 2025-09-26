package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestGetRelatedDocsHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := Site{
		relatedDocsHandler: &page.RelatedDocsHandler{},
	}

	handler := s.GetRelatedDocsHandler()

	if handler != s.relatedDocsHandler {
		t.Errorf("Expected %v, got %v", s.relatedDocsHandler, handler)
	}
}
