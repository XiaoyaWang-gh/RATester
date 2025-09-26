package legacy

import (
	"fmt"
	"testing"
)

func TestConvert_ClientCommonConf_To_v1_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := &ClientCommonConf{
		User:       "test_user",
		ServerAddr: "127.0.0.1",
		// other fields...
	}
	out := Convert_ClientCommonConf_To_v1(conf)

	// assertions
	if out.User != conf.User {
		t.Errorf("Expected User to be %s, but got %s", conf.User, out.User)
	}
	if out.ServerAddr != conf.ServerAddr {
		t.Errorf("Expected ServerAddr to be %s, but got %s", conf.ServerAddr, out.ServerAddr)
	}
	// add more assertions for other fields...
}
