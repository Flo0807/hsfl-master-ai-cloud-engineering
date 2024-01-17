package balancer

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/load-balancer/model"
	"github.com/stretchr/testify/assert"
)

type mockAlgorithm struct {
	getTargetCalled bool
}

func (m *mockAlgorithm) GetTarget(r *http.Request, targets []model.Target, callback func(target model.Target)) {
	m.getTargetCalled = true
	callback(targets[0])
}

func TestNewBalancer(t *testing.T) {
	// given
	a := &mockAlgorithm{}
	targets := []model.Target{{}, {}, {}}

	// when
	b := NewBalancer(a, targets, 10, "/health")

	// test
	assert.Equal(t, a, b.algorithm)
	assert.Equal(t, targets, b.targets)
	assert.Equal(t, targets, b.healthyTargets)
	assert.Equal(t, 10, b.healthCheckInterval)
}

func TestServeHTTP(t *testing.T) {
	// given
	a := &mockAlgorithm{}
	targetURL, _ := url.Parse("http://localhost:8080")
	targets := []model.Target{{Url: targetURL}, {}, {}}

	// when
	b := NewBalancer(a, targets, 10, "/health")
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	w := httptest.NewRecorder()

	b.ServeHTTP(w, req)

	// test
	assert.True(t, a.getTargetCalled)
}

func TestHealthCheck(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	targetURL, _ := url.Parse(server.URL)
	targets := []model.Target{{Url: targetURL}}
	b := NewBalancer(&mockAlgorithm{}, targets, 1, "/health")

	time.Sleep(2 * time.Second)

	assert.Contains(t, b.healthyTargets, targets[0])

	server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	time.Sleep(2 * time.Second)

	assert.NotContains(t, b.healthyTargets, targets[0])
}
