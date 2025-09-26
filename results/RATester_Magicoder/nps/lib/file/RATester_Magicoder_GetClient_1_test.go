package file

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGetClient_1(t *testing.T) {
	db := &DbUtils{
		JsonDb: &JsonDb{
			Clients: sync.Map{},
		},
	}

	client := &Client{
		Id: 1,
	}

	db.JsonDb.Clients.Store(client.Id, client)

	testCases := []struct {
		name     string
		id       int
		expected *Client
		err      error
	}{
		{
			name:     "client exists",
			id:       1,
			expected: client,
			err:      nil,
		},
		{
			name:     "client not exists",
			id:       2,
			expected: nil,
			err:      errors.New("未找到客户端"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := db.GetClient(tc.id)
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if err == nil && !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
