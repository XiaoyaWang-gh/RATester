package sub

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/config"
)

func TestExecute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rootCmd.SetGlobalNormalizationFunc(config.WordSepNormalizeFunc)
	if err := rootCmd.Execute(); err != nil {
		t.Errorf("Execute() error = %v", err)
	}
}
