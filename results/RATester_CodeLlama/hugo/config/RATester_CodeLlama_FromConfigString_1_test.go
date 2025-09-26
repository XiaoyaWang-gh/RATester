package config

import (
	"fmt"
	"testing"
)

func TestFromConfigString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var (
		configType = "yaml"
		config     = `
title: "My Site"
baseURL: "https://example.com"
`
	)

	p, err := FromConfigString(config, configType)
	if err != nil {
		t.Fatal(err)
	}

	if p.GetString("title") != "My Site" {
		t.Errorf("title should be My Site, got %q", p.GetString("title"))
	}

	if p.GetString("baseURL") != "https://example.com" {
		t.Errorf("baseURL should be https://example.com, got %q", p.GetString("baseURL"))
	}
}
