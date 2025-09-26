package org

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestNew_56(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := provide{}
	cfg := converter.ProviderConfig{}
	_, err := p.New(cfg)
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
}
