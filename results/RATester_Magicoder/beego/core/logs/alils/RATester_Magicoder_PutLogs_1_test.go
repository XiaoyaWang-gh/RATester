package alils

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestPutLogs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &LogGroup{
		Logs: []*Log{
			{
				Time: proto.Uint32(1609459200),
				Contents: []*LogContent{
					{
						Key:   proto.String("test_key"),
						Value: proto.String("test_value"),
					},
				},
			},
		},
	}

	ls := &LogStore{
		Name: "test_logstore",
		project: &LogProject{
			Name:            "test_project",
			Endpoint:        "http://test_endpoint",
			AccessKeyID:     "test_access_key_id",
			AccessKeySecret: "test_access_key_secret",
		},
	}

	err := ls.PutLogs(lg)
	if err != nil {
		t.Errorf("PutLogs failed: %v", err)
	}
}
