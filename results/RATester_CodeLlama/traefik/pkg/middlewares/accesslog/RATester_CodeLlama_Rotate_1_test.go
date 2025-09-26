package accesslog

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestRotate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Handler{}
	h.config = &types.AccessLog{}
	h.config.FilePath = "test.log"
	h.logger = logrus.New()
	h.logger.Out = os.Stdout
	h.mu = sync.Mutex{}
	h.httpCodeRanges = types.HTTPCodeRanges{}
	h.logHandlerChan = make(chan handlerParams, 100)
	h.wg = sync.WaitGroup{}

	err := h.Rotate()
	require.NoError(t, err)
}
