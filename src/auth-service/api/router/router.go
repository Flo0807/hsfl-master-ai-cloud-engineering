package router

import (
	"net/http"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/router"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/auth-service/api/handler"
)

type Router struct {
	router http.Handler
}

func NewRouter(loginHandler *handler.LoginHandler, registerHandler *handler.RegisterHandler) *Router {
	r := router.New()

	r.POST("/login", loginHandler.Login)
	r.POST("/register", registerHandler.Register)

	return &Router{r}
}

func (handler *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
