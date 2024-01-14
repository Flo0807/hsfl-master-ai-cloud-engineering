package router

import (
	"net/http"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/router"
	middleware "github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/router/middleware/bulletinboard"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/feed-service/feed"
)

type Router struct {
	router http.Handler
}

func New(feedController *feed.DefaultController, bulletinBoardServiceClient bulletin_board.BulletinBoardServiceClient) *Router {
	bulletinBoardMiddleware := middleware.CreateBulletinBoardMiddleware(bulletinBoardServiceClient)

	r := router.New()

	r.GET("/feed/feed", feedController.GetFeed, bulletinBoardMiddleware)
	r.GET("/feed/health", feedController.GetHealth)

	return &Router{r}
}

func (handler *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
