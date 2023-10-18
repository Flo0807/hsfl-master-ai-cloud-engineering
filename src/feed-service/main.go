package main
import (
	"net/http"
	"log"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/feed-service/feed"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/feed-service/api/router"

)
func main(){
	feedController := feed.NewDefaultController()
	handler := router.New(feedController)
	if err := http.ListenAndServe(":3000", handler); err != nil {
		log.Fatalf("error while listen and serve: %s", err.Error())
	}
}