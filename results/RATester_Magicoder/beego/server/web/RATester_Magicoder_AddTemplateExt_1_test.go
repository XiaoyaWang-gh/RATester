package web

import (
	"fmt"
	"testing"
)

func TestAddTemplateExt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	beeTemplateExt = []string{}
	AddTemplateExt("tpl")
	if len(beeTemplateExt) != 1 {
		t.Errorf("Expected length of beeTemplateExt to be 1, but got %d", len(beeTemplateExt))
	}
	if beeTemplateExt[0] != "tpl" {
		t.Errorf("Expected first element of beeTemplateExt to be 'tpl', but got '%s'", beeTemplateExt[0])
	}

	AddTemplateExt("tpl")
	if len(beeTemplateExt) != 1 {
		t.Errorf("Expected length of beeTemplateExt to be 1, but got %d", len(beeTemplateExt))
	}
	if beeTemplateExt[0] != "tpl" {
		t.Errorf("Expected first element of beeTemplateExt to be 'tpl', but got '%s'", beeTemplateExt[0])
	}

	AddTemplateExt("html")
	if len(beeTemplateExt) != 2 {
		t.Errorf("Expected length of beeTemplateExt to be 2, but got %d", len(beeTemplateExt))
	}
	if beeTemplateExt[1] != "html" {
		t.Errorf("Expected second element of beeTemplateExt to be 'html', but got '%s'", beeTemplateExt[1])
	}
}
