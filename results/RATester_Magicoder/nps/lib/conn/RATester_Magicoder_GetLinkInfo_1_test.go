package conn

import (
	"fmt"
	"testing"
)

func TestGetLinkInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: nil,
		Rb:   []byte{},
	}

	link, err := conn.GetLinkInfo()
	if err != nil {
		t.Errorf("GetLinkInfo() error = %v", err)
		return
	}

	if link == nil {
		t.Errorf("GetLinkInfo() link is nil")
		return
	}

	if link.ConnType != "" {
		t.Errorf("GetLinkInfo() link.ConnType = %v, want %v", link.ConnType, "")
	}

	if link.Host != "" {
		t.Errorf("GetLinkInfo() link.Host = %v, want %v", link.Host, "")
	}

	if link.Crypt != false {
		t.Errorf("GetLinkInfo() link.Crypt = %v, want %v", link.Crypt, false)
	}

	if link.Compress != false {
		t.Errorf("GetLinkInfo() link.Compress = %v, want %v", link.Compress, false)
	}

	if link.LocalProxy != false {
		t.Errorf("GetLinkInfo() link.LocalProxy = %v, want %v", link.LocalProxy, false)
	}

	if link.RemoteAddr != "" {
		t.Errorf("GetLinkInfo() link.RemoteAddr = %v, want %v", link.RemoteAddr, "")
	}
}
