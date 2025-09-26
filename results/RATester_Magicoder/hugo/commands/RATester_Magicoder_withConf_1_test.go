package commands

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/config/allconfig"
)

func TestwithConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hb := &hugoBuilder{
		conf: &commonConfig{
			mu:      &sync.Mutex{},
			configs: &allconfig.Configs{},
		},
	}

	hb.withConf(func(conf *commonConfig) {
		conf.mu.Lock()
		defer conf.mu.Unlock()
		// Add your test logic here
	})
}
