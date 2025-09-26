package install

import (
	"fmt"
	"os"
	"testing"

	"ehang.io/nps/lib/common"
)

func TestInstallNpc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := common.GetInstallPath()
	if !common.FileExists(path) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			t.Fatal(err)
		}
	}
	copyStaticFile(common.GetAppPath(), "npc")
}
