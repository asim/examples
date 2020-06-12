package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"

	tags "github.com/micro/examples/blog/tags/proto/tags"

	"github.com/gosimple/slug"
)

const (
	parentPrefix = "parent"
	typePrefix   = "type"
)

type Tag struct {
	ParentID string `json:"parentID"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Type     string `json:"type"`
	Count    int64  `json:"count"`
}

type Tags struct {
	Store store.Store
}

func (t *Tags) IncreaseCount(ctx context.Context, req *tags.IncreaseCountRequest, rsp *tags.IncreaseCountResponse) error {
	if len(req.ParentID) == 0 || len(req.Type) == 0 {
		return errors.New("parent id and type is required")
	}

	tagSlug := slug.Make(req.GetTitle())
	parentID := fmt.Sprintf("%v:%v:%v", parentPrefix, req.GetParentID(), tagSlug)

	// read by parent ID + slug, the record is identical in boths places anyway
	records, err := t.Store.Read(parentID)
	if err != nil && err != store.ErrNotFound {
		return err
	}

	// If no existing record is found, create a new one
	if len(records) == 0 {
		tag := &Tag{
			ParentID: req.GetParentID(),
			Title:    req.GetTitle(),
			Type:     req.Type,
			Slug:     tagSlug,
			Count:    1,
		}
		return t.saveTag(tag)
	}
	record := records[0]
	tag := &Tag{}
	err = json.Unmarshal(record.Value, tag)
	if err != nil {
		return err
	}
	tag.Count++
	return t.saveTag(tag)
}

func (t *Tags) saveTag(tag *Tag) error {
	tagSlug := slug.Make(tag.Title)

	parentID := fmt.Sprintf("%v:%v:%v", parentPrefix, tag.ParentID, tagSlug)
	typeID := fmt.Sprintf("%v:%v:%v", typePrefix, tag.Type, tagSlug)

	bytes, err := json.Marshal(tag)
	if err != nil {
		return err
	}

	// write parentId:slug to enable prefix listing based on parent
	err = t.Store.Write(&store.Record{
		Key:   parentID,
		Value: bytes,
	})
	if err != nil {
		return err
	}

	// write type:slug to enable prefix listing based on parent
	return t.Store.Write(&store.Record{
		Key:   typeID,
		Value: bytes,
	})
}

func (t *Tags) DecreaseCount(ctx context.Context, req *tags.DecreaseCountRequest, rsp *tags.DecreaseCountResponse) error {
	if len(req.ParentID) == 0 || len(req.Type) == 0 {
		return errors.New("parent id and type is required")
	}

	tagSlug := slug.Make(req.GetTitle())
	parentID := fmt.Sprintf("%v:%v:%v", parentPrefix, req.GetParentID(), tagSlug)

	// read by parent ID + slug, the record is identical in boths places anyway
	records, err := t.Store.Read(parentID)
	if err != nil && err != store.ErrNotFound {
		return err
	}

	// If no existing record is found, there is nothing to decrease
	if len(records) == 0 {
		// return error?
		return nil
	}
	record := records[0]
	tag := &Tag{}
	err = json.Unmarshal(record.Value, tag)
	if err != nil {
		return err
	}
	if tag.Count == 0 {
		// return error?
		return nil
	}
	tag.Count--
	return t.saveTag(tag)
}

func (t *Tags) List(ctx context.Context, req *tags.ListRequest, rsp *tags.ListResponse) error {
	log.Info("Received tags.List request")
	key := ""
	if len(req.ParentID) > 0 {
		key = fmt.Sprintf("%v:%v", parentPrefix, req.ParentID)
	} else if len(req.Type) > 0 {
		key = fmt.Sprintf("%v:%v", typePrefix, req.Type)
	} else {
		return errors.New("parent id or type required for listing")
	}

	records, err := t.Store.Read(key, store.ReadPrefix())
	if err != nil {
		return err
	}
	rsp.Tags = make([]*tags.Tag, len(records))
	for i, record := range records {
		tagRecord := &Tag{}
		err := json.Unmarshal(record.Value, tagRecord)
		if err != nil {
			return err
		}
		rsp.Tags[i] = &tags.Tag{
			ParentID: tagRecord.ParentID,
			Title:    tagRecord.Title,
			Type:     tagRecord.Type,
			Slug:     tagRecord.Slug,
			Count:    tagRecord.Count,
		}
	}
	return nil
}

func (t *Tags) Update(ctx context.Context, req *tags.UpdateRequest, rsp *tags.UpdateResponse) error {
	if len(req.ParentID) == 0 || len(req.Type) == 0 {
		return errors.New("parent id and type is required")
	}

	tagSlug := slug.Make(req.GetTitle())
	parentID := fmt.Sprintf("%v:%v:%v", parentPrefix, req.GetParentID(), tagSlug)

	// read by parent ID + slug, the record is identical in boths places anyway
	records, err := t.Store.Read(parentID)
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("Tag with slug '%v' not found, nothing to update", tagSlug)
	}
	record := records[0]
	tag := &Tag{}
	err = json.Unmarshal(record.Value, tag)
	if err != nil {
		return err
	}
	tag.Title = req.Title
	return t.saveTag(tag)
}
