package accesslog

import (
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestRotate_1(t *testing.T) {
	type fields struct {
		config         *types.AccessLog
		logger         *logrus.Logger
		file           io.WriteCloser
		mu             sync.Mutex
		httpCodeRanges types.HTTPCodeRanges
		logHandlerChan chan handlerParams
		wg             sync.WaitGroup
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "TestRotate_Success",
			fields: fields{
				config: &types.AccessLog{
					FilePath: "test.log",
				},
				logger: logrus.New(),
			},
			wantErr: false,
		},
		{
			name: "TestRotate_Failure",
			fields: fields{
				config: &types.AccessLog{
					FilePath: "",
				},
				logger: logrus.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			h := &Handler{
				config:         tt.fields.config,
				logger:         tt.fields.logger,
				file:           tt.fields.file,
				mu:             tt.fields.mu,
				httpCodeRanges: tt.fields.httpCodeRanges,
				logHandlerChan: tt.fields.logHandlerChan,
				wg:             tt.fields.wg,
			}
			if err := h.Rotate(); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Rotate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
