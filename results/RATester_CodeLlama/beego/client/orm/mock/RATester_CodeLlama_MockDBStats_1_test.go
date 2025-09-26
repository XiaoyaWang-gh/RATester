package mock

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMockDBStats_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given:
	stats := &sql.DBStats{}
	// when:
	mock := MockDBStats(stats)
	// then:
	assert.Equal(t, stats, mock.resp[0])
}
