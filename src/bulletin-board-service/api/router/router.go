package router

import (
	"net/http"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/bulletin-board-service/api/handler"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/router"
	middleware "github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/router/middleware/auth"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/auth"
)

type Router struct {
	router http.Handler
}

func NewRouter(
	healthHandler *handler.HealthHandler,
	postHandler *handler.PostHandler,
	authServiceClient auth.AuthServiceClient,
) *Router {
	authMiddleware := middleware.CreateAuthMiddleware(authServiceClient)

	r := router.New()
	r.GET("/bulletin-board/health", healthHandler.Health)
	r.GET("/bulletin-board/posts", postHandler.GetPosts)
	r.GET("/bulletin-board/posts-request-coalescing", postHandler.GetPostsRequestCoalescing)
	r.GET("/bulletin-board/posts/:id", postHandler.GetPost)
	r.POST("/bulletin-board/posts", postHandler.CreatePost)
	r.PUT("/bulletin-board/posts/:id", postHandler.UpdatePost)
	r.DELETE("/bulletin-board/posts/:id", postHandler.DeletePost, authMiddleware)

	return &Router{r}
}

func (handler *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
