# likit-go

`likit-go` is a Client library for the Likit in Golang.


## Usage
```
go get github.com/likit/likit-go
```

A Comment Like Example

```go
import (
	"github.com/CorrectRoadH/likit-go"
)

type CommentServer struct {
	Store      *store.CommentStore
	likitServer *likit.LikitServer
}

func NewCommentServer(store *store.CommentStore) *CommentServer {
	return &CommentServer{
		Store:      store,
        // the likit server address. You should replace it with your own server address
		// likit.NewlikitServer("localhost:4778",false), 
		likitServer: likit.NewlikitServer("likit.zeabur.app:443",true), 
		// NewlikitServer(host string, tls bool)
		// if you deploy likit server in zeabur. tls shoud be true
	}
}

const businessId = "COMMENT_LIKE"
// get comment
func (s *CommentServer) ListComment(c echo.Context)  error {
    userId := /// get user id from jwt or cookie

    // define find condition
    find := ////
	comments, err := s.Store.ListComment(ctx, find)

	return c.JSON(http.StatusOK, lo.Map(comments, func(comment types.Comment, index int) Comment {
        isLike, err := s.likitServer.IsVote(ctx, businessId, comment.Id, comment.UserId)
        if err != nil {
            // handle error
        }

        likeNum, err := s.likitServer.Count(ctx, businessId, comment.Id)
        if err != nil {
            // handle err
        }

        return convertComment(comment, isLike, likeNum)
    }))
}

func (s *CommentServer) Like(c echo.Context) (*apiv1pb.LikeResponse, error) {
	userId := /// get user id from jwt or cookie
	commentId := /// get comment id from request

	count, err := s.likitServer.Vote(ctx, businessId, commentId, UserId)
	if err != nil {
		// handle error
	}
	return c.JSON(http.StatusOK, Count: count)
}

func (s *CommentServer) Unlike(c echo.Context) (*apiv1pb.LikeResponse, error) {
    userId := /// get user id from jwt or cookie
	commentId := /// get comment id from request

	count, err := s.likitServer.UnVote(ctx, businessId, commentId, UserId)
	if err != nil {
		// handle error
	}

	return c.JSON(http.StatusOK, Count: count)
}
```