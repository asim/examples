package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"

	tagProto "github.com/micro/examples/blog/tags/proto/tags"

	"github.com/gosimple/slug"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"

	posts "github.com/micro/examples/blog/posts/proto/posts"
)

const (
	tagType         = "post-tag"
	slugPrefix      = "slug"
	idPrefix        = "id"
	timeStampPrefix = "timestamp"
)

type Post struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Slug            string   `json:"slug"`
	Content         string   `json:"content"`
	CreateTimestamp int64    `json:"create_timestamp"`
	UpdateTimestamp int64    `json:"update_timestamp"`
	TagNames        []string `json:"tagNames"`
}

type Posts struct {
	Store  store.Store
	Client client.Client
}

func (t *Posts) Post(ctx context.Context, req *posts.PostRequest, rsp *posts.PostResponse) error {
	if len(req.Post.Id) == 0 || len(req.Post.Title) == 0 || len(req.Post.Content) == 0 {
		return errors.New("ID, title or content is missing")
	}

	// read by post
	records, err := t.Store.Read(fmt.Sprintf("%v:%v", idPrefix, req.Post.Id))
	if err != nil && err != store.ErrNotFound {
		return err
	}
	postSlug := slug.Make(req.Post.Title)
	// If no existing record is found, create a new one
	if len(records) == 0 {
		post := &Post{
			ID:              req.Post.Id,
			Title:           req.Post.Title,
			Content:         req.Post.Content,
			TagNames:        req.Post.TagNames,
			Slug:            postSlug,
			CreateTimestamp: time.Now().Unix(),
		}
		return t.savePost(ctx, nil, post)
	}
	record := records[0]
	oldPost := &Post{}
	err = json.Unmarshal(record.Value, oldPost)
	if err != nil {
		return err
	}
	post := &Post{
		ID:              req.Post.Id,
		Title:           req.Post.Title,
		Content:         req.Post.Content,
		Slug:            postSlug,
		TagNames:        req.Post.TagNames,
		CreateTimestamp: oldPost.CreateTimestamp,
		UpdateTimestamp: time.Now().Unix(),
	}

	// Check if slug exists
	recordsBySlug, err := t.Store.Read(fmt.Sprintf("%v:%v", slugPrefix, postSlug))
	if err != nil && err != store.ErrNotFound {
		return err
	}
	otherSlugPost := &Post{}
	err = json.Unmarshal(record.Value, otherSlugPost)
	if err != nil {
		return err
	}
	if len(recordsBySlug) > 0 && oldPost.ID != otherSlugPost.ID {
		return errors.New("An other post with this slug already exists")
	}

	return t.savePost(ctx, oldPost, post)
}

func (t *Posts) savePost(ctx context.Context, oldPost, post *Post) error {
	bytes, err := json.Marshal(post)
	if err != nil {
		return err
	}

	err = t.Store.Write(&store.Record{
		Key:   fmt.Sprintf("%v:%v", idPrefix, post.ID),
		Value: bytes,
	})
	if err != nil {
		return err
	}
	// Delete old slug index if the slug has changed
	if oldPost.Slug != post.Slug {
		err = t.Store.Delete(fmt.Sprintf("%v:%v", slugPrefix, post.Slug))
		if err != nil {
			return err
		}
	}
	err = t.Store.Write(&store.Record{
		Key:   fmt.Sprintf("%v:%v", slugPrefix, post.Slug),
		Value: bytes,
	})
	if err != nil {
		return err
	}
	err = t.Store.Write(&store.Record{
		Key:   fmt.Sprintf("%v:%v", timeStampPrefix, math.MaxInt64-post.CreateTimestamp),
		Value: bytes,
	})
	if err != nil {
		return err
	}
	if oldPost == nil {
		tagClient := tagProto.NewTagsService("go.micro.service.tag", t.Client)
		for _, tagName := range post.TagNames {
			_, err := tagClient.IncreaseCount(ctx, &tagProto.IncreaseCountRequest{
				ParentID: post.ID,
				Type:     tagType,
				Title:    tagName,
			})
			if err != nil {
				return err
			}
		}
		return nil
	}
	return t.diffTags(ctx, oldPost.ID, oldPost.TagNames, post.TagNames)
}

func (t *Posts) diffTags(ctx context.Context, parentID string, oldTagNames, newTagNames []string) error {
	oldTags := map[string]struct{}{}
	for _, v := range oldTagNames {
		oldTags[v] = struct{}{}
	}
	newTags := map[string]struct{}{}
	for _, v := range newTagNames {
		newTags[v] = struct{}{}
	}
	tagClient := tagProto.NewTagsService("go.micro.service.tag", t.Client)
	for i := range oldTags {
		_, stillThere := newTags[i]
		if !stillThere {
			tagClient.DecreaseCount(ctx, &tagProto.DecreaseCountRequest{
				ParentID: parentID,
				Type:     tagType,
				Title:    i,
			})
		}
	}
	for i := range newTags {
		_, newlyAdded := oldTags[i]
		if newlyAdded {
			tagClient.IncreaseCount(ctx, &tagProto.IncreaseCountRequest{
				ParentID: parentID,
				Type:     tagType,
				Title:    i,
			})
		}
	}
	return nil
}

func (t *Posts) Query(ctx context.Context, req *posts.QueryRequest, rsp *posts.QueryResponse) error {
	var records []*store.Record
	var err error
	if len(req.Slug) > 0 {
		key := fmt.Sprintf("%v:%v", slugPrefix, req.Slug)
		log.Infof("Reading post by slug: %v", req.Slug)
		records, err = t.Store.Read(key, store.ReadPrefix())
	} else {
		key := fmt.Sprintf("%v:", timeStampPrefix)
		var limit uint
		limit = 20
		if req.Limit > 0 {
			limit = uint(req.Limit)
		}
		log.Infof("Listing posts, offset: %v, limit: %v", req.Offset, limit)
		records, err = t.Store.Read(key, store.ReadPrefix(),
			store.ReadOffset(uint(req.Offset)),
			store.ReadLimit(limit))
	}

	if err != nil {
		return err
	}
	rsp.Posts = make([]*posts.Post, len(records))
	for i, record := range records {
		postRecord := &Post{}
		err := json.Unmarshal(record.Value, postRecord)
		if err != nil {
			return err
		}
		rsp.Posts[i] = &posts.Post{
			Id:       postRecord.ID,
			Title:    postRecord.Title,
			Slug:     postRecord.Slug,
			Content:  postRecord.Content,
			TagNames: postRecord.TagNames,
		}
	}
	return nil
}

func (t *Posts) Delete(ctx context.Context, req *posts.DeleteRequest, rsp *posts.DeleteResponse) error {
	log.Info("Received Post.Delete request")
	records, err := t.Store.Read(fmt.Sprintf("%v:%v", idPrefix, req.Id))
	if err != nil && err != store.ErrNotFound {
		return err
	}
	if len(records) == 0 {
		return fmt.Errorf("Post with ID %v not found", req.Id)
	}
	post := &Post{}
	err = json.Unmarshal(records[0].Value, post)
	if err != nil {
		return err
	}

	// Delete by ID
	err = t.Store.Delete(fmt.Sprintf("%v:%v", idPrefix, post.ID))
	if err != nil {
		return err
	}
	// Delete by slug
	err = t.Store.Delete(fmt.Sprintf("%v:%v", slugPrefix, post.Slug))
	if err != nil {
		return err
	}
	// Delete by timeStamp
	return t.Store.Delete(fmt.Sprintf("%v:%v", timeStampPrefix, post.CreateTimestamp))
}
