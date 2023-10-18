package feed

import(
	"net/http"
)

type Controller interface{
	GetFeed(http.ResponseWriter, *http.Request)
}