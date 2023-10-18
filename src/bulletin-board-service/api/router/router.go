package router

import (
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/api/handler"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/router"
	"net/http"
)

type Router struct {
	router http.Handler
}

func NewRouter(postHandler *handler.PostHandler) *Router {
	r := router.New()

	r.GET("/posts", postHandler.GetPosts)
	r.GET("/posts/:id", postHandler.GetPost)
	r.POST("/posts", postHandler.CreatePost)
	r.PUT("/posts/:id", postHandler.UpdatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)

	return &Router{r}
}

func (handler *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
