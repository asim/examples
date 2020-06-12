package subscriber

import (
	"context"

	post "github.com/micro/examples/blog/post/proto/post"
)

type Post struct{}

func (e *Post) Handle(ctx context.Context, msg *post.Post) error {
	return nil
}

func Handler(ctx context.Context, msg *post.Post) error {
	return nil
}
