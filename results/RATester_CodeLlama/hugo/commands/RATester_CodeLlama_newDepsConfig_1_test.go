package commands

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/deps"
)

func TestNewDepsConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	conf := &commonConfig{}
	r := &rootCommand{}
	// When
	result := r.newDepsConfig(conf)
	// Then
	assert.Equal(t, deps.DepsCfg{Configs: conf.configs, Fs: conf.fs, LogOut: r.logger.Out(), LogLevel: r.logger.Level(), ChangesFromBuild: r.changesFromBuild}, result)
}
