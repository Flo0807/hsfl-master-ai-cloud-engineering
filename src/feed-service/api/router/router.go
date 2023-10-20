package router

import (
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/router"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/feed-service/feed"
	"net/http"
)

type Router struct {
	router http.Handler
}

func New(feedController *feed.DefaultController) *Router {
	r := router.New()

	r.GET("/feed", feedController.GetFeed)

	return &Router{r}
}

func (handler *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
