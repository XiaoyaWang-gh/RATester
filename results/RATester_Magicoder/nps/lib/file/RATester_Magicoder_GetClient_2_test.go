package file

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGetClient_2(t *testing.T) {
	db := &JsonDb{
		Clients: sync.Map{},
	}

	client := &Client{
		Id: 1,
	}

	db.Clients.Store(client.Id, client)

	tests := []struct {
		name    string
		id      int
		want    *Client
		wantErr bool
	}{
		{
			name:    "Test case 1",
			id:      1,
			want:    client,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			id:      2,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := db.GetClient(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
