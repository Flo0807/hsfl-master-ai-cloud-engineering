package least_connections

import (
	"math"
	"net/http"
	"sync"
)

type LeastConnections struct {
	mutex sync.RWMutex
	count map[string]int
}

func New() *LeastConnections {
	return &LeastConnections{sync.RWMutex{}, make(map[string]int)}
}

func (algorithm *LeastConnections) GetTarget(_ *http.Request, replicas []string, fun func(host string)) {
	algorithm.mutex.RLock()

	var targetHost string

	minConnections := math.MaxInt

	for _, host := range replicas {
		if currentConnections, ok := algorithm.count[host]; ok && currentConnections < minConnections {
			minConnections = currentConnections
			targetHost = host
		}
	}

	algorithm.mutex.RUnlock()

	algorithm.adjustConnectionCount(targetHost, 1)

	fun(targetHost)

	algorithm.adjustConnectionCount(targetHost, -1)
}

func (algorithm *LeastConnections) adjustConnectionCount(host string, delta int) {
	algorithm.mutex.Lock()
	defer algorithm.mutex.Unlock()

	algorithm.count[host] += delta
}
