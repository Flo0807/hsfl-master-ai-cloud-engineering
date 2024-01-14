package router

import (
	"net/http"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/router"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/feed-service/feed"
)

type Router struct {
	router http.Handler
}

func New(feedController *feed.DefaultController) *Router {
	r := router.New()

	r.GET("/feed/feed", feedController.GetFeed)
	r.GET("/feed/health", feedController.GetHealth)

	return &Router{r}
}

func (handler *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
