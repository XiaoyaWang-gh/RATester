package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestLoadConfigMain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &configLoader{}
	d := ConfigSourceDescriptor{}
	res, _, err := l.loadConfigMain(d)
	if err != nil {
		t.Errorf("loadConfigMain failed: %s", err)
	}
	if res.Cfg == nil {
		t.Errorf("loadConfigMain failed: res.Cfg is nil")
	}
	if res.ConfigFiles == nil {
		t.Errorf("loadConfigMain failed: res.ConfigFiles is nil")
	}
	if res.BaseConfig == (config.BaseConfig{}) {
		t.Errorf("loadConfigMain failed: res.BaseConfig is empty")
	}
}
