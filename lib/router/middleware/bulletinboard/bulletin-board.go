package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/router"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/lib/rpc/bulletin-board/rpc/bulletin_board"
)

type UserIdKey string

func CreateBulletinBoardMiddleware(bulletinBoardServiceClient bulletin_board.BulletinBoardServiceClient) router.MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("In Middleware")
			resp, err := bulletinBoardServiceClient.GetPosts(context.Background(), &bulletin_board.Request{})
			for _, v := range resp.Posts {
				fmt.Println(v.Content)

			}
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "posts", resp.Posts)
			r = r.WithContext(ctx)

			next(w, r)
		}
	}
}
