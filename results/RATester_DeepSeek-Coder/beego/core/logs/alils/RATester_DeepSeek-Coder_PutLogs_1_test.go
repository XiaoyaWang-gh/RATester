package alils

import (
	"fmt"
	"testing"
	"time"

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
				Time: proto.Uint32(uint32(time.Now().Unix())),
				Contents: []*LogContent{
					{
						Key:   proto.String("test_key"),
						Value: proto.String("test_value"),
					},
				},
			},
		},
		Topic:  proto.String("test_topic"),
		Source: proto.String("test_source"),
	}

	s := &LogStore{
		Name: "test_logstore",
		project: &LogProject{
			Name:            "test_project",
			Endpoint:        "http://localhost:8080",
			AccessKeyID:     "test_access_key_id",
			AccessKeySecret: "test_access_key_secret",
		},
	}

	err := s.PutLogs(lg)
	if err != nil {
		t.Errorf("PutLogs failed: %v", err)
	}
}
