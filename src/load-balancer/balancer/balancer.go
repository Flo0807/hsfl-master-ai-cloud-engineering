package balancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/load-balancer/algorithm"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/load-balancer/model"
)

type Balancer struct {
	algorithm           algorithm.Algorithm
	targets             []model.Target
	healthyTargets      []model.Target
	healthCheckInterval int
	healthCheckPath     string
}

func NewBalancer(algorithm algorithm.Algorithm, targets []model.Target, healthCheckInterval int, healthCheckPath string) *Balancer {
	b := &Balancer{algorithm: algorithm, targets: targets, healthyTargets: targets, healthCheckInterval: healthCheckInterval, healthCheckPath: healthCheckPath}
	go b.healthCheck()
	return b
}

func (b *Balancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.algorithm.GetTarget(r, b.healthyTargets, func(target model.Target) {
		httputil.NewSingleHostReverseProxy(target.Url).ServeHTTP(w, r)
	})
}

func (b *Balancer) healthCheck() {
	for {
		for i, target := range b.targets {
			if target.Url == nil {
				continue
			}

			resp, err := http.Get(target.Url.String() + b.healthCheckPath)
			if err != nil || resp.StatusCode != http.StatusOK {
				// Remove from healthyTargets if it's there
				for j, healthyTarget := range b.healthyTargets {
					if target == healthyTarget {
						b.healthyTargets = append(b.healthyTargets[:j], b.healthyTargets[j+1:]...)
						break
					}
				}

				log.Printf("Target %s is unhealthy", target.Url.String())
			} else {
				// Add to healthyTargets if it's not there
				found := false
				for _, healthyTarget := range b.healthyTargets {
					if target == healthyTarget {
						found = true
						break
					}
				}
				if !found {
					b.healthyTargets = append(b.healthyTargets, b.targets[i])

					log.Printf("Target %s is healthy again", target.Url.String())
				}
			}
		}
		time.Sleep(time.Duration(b.healthCheckInterval) * time.Second)
	}
}
