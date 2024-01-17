package least_connections

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastConnections(t *testing.T) {
	// given
	lc := New()

	lc.adjustConnectionCount("host1", 5)
	lc.adjustConnectionCount("host2", 3)
	lc.adjustConnectionCount("host3", 7)

	var targetHost string

	// when
	lc.GetTarget(&http.Request{}, []string{"host1", "host2", "host3"}, func(host string) {
		targetHost = host
	})

	// test
	assert.Equal(t, "host2", targetHost)
}

func TestAdjustConnectionCount(t *testing.T) {
	// given
	lc := New()

	// when
	lc.adjustConnectionCount("host1", 1)
	lc.adjustConnectionCount("host2", -1)
	lc.adjustConnectionCount("host3", 3)

	// test
	assert.Equal(t, 1, lc.count["host1"])
	assert.Equal(t, -1, lc.count["host2"])
	assert.Equal(t, 3, lc.count["host3"])
}
