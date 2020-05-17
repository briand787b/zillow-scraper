package redisclient

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"zcrapr/core/model"
	"zcrapr/core/perr"
	"zcrapr/core/plog"

	goredis "github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

// PropertyRedisStore is a PropertyStore backed by Redis
type PropertyRedisStore struct {
	l      plog.Logger
	client *goredis.Client

	idCounterKey  string
	idPrefix      string
	captureSuffix string
	idWidth       string
}

// NewPropertyRedisStore returns a new PropertyStore backed by Redis
func NewPropertyRedisStore(l plog.Logger, idCounterKey, idPrefix, captureSuffix, host, password string, idWidth, port uint) (*PropertyRedisStore, error) {
	if idCounterKey == "" {
		return nil, perr.NewErrInvalid("idCounterKey cannot be empty string")
	}

	if idPrefix == "" {
		return nil, perr.NewErrInvalid("idPrefix cannot be empty string")
	}

	if captureSuffix == "" {
		return nil, perr.NewErrInvalid("captureSuffix cannot be empty string")
	}

	if idWidth == 0 {
		return nil, perr.NewErrInvalid("idWidth cannot be zero")
	}

	client := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%v", host, port),
		Password: password,
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		l.Error(nil, pong, "error", err)
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not ping Redis")
	}

	return &PropertyRedisStore{
		l:             l,
		client:        client,
		idCounterKey:  idCounterKey,
		idPrefix:      idPrefix,
		captureSuffix: captureSuffix,
		idWidth:       strconv.Itoa(int(idWidth)),
	}, nil
}

// GetAllCapturesByPropertyID x
func (s *PropertyRedisStore) GetAllCapturesByPropertyID(ctx context.Context, propID string) ([]model.Capture, error) {
	captureKey := propID + s.captureSuffix
	caps, err := s.client.LRange(captureKey, 0, -1).Result()
	if err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not range over captures")
	}

	modelCaps := make([]model.Capture, len(caps))
	for i := 0; i < len(caps); i++ {
		if err := json.Unmarshal([]byte(caps[i]), &modelCaps[i]); err != nil {
			return nil, errors.Wrap(perr.NewErrInternal(err), "could not unmarshal Capture to JSON")
		}
	}

	return modelCaps, nil
}

// GetAllPropertyIDs x
func (s *PropertyRedisStore) GetAllPropertyIDs(ctx context.Context, skip, take int) ([]string, error) {
	var allKeys []string
	var nextKeys []string
	var err error
	takeIncr := int64(100)
	cursor := uint64(skip)
	remaining := int64(take)
	for remaining > 0 {
		if takeIncr > remaining {
			takeIncr = remaining
		}

		nextKeys, cursor, err = s.client.Scan(cursor, s.idPrefix+":*", takeIncr).Result()
		if err != nil {
			return nil, perr.NewErrInternal(errors.Wrap(err, "could not execute Redis command"))
		}

		allKeys = append(allKeys, nextKeys...)
		remaining -= takeIncr
	}

	return allKeys, nil
}

// GetLatestCaptureByPropertyID x
func (s *PropertyRedisStore) GetLatestCaptureByPropertyID(ctx context.Context, propID string) (*model.Capture, error) {
	captureKey := propID + s.captureSuffix
	caps, err := s.client.LRange(captureKey, 0, 0).Result()
	if err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not execute Redis command")
	}

	if len(caps) < 1 {
		return nil, perr.NewErrNotFound(errors.New("no captures found"))
	}

	var cap model.Capture
	if err := json.Unmarshal([]byte(caps[0]), &cap); err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not unmarshal to Redis output to Capture")
	}

	return &cap, nil
}

// GetPropertyByID x
func (s *PropertyRedisStore) GetPropertyByID(ctx context.Context, id string) (*model.Property, error) {
	fieldMap, err := s.client.HGetAll(id).Result()
	if err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not get all hash keys")
	}

	id, ok := fieldMap["id"]
	if !ok || id == "" {
		return nil, perr.NewErrInternal(errors.New("returned ID is empty"))
	}

	url, ok := fieldMap["url"]
	if !ok || id == "" {
		return nil, perr.NewErrInternal(errors.New("returned url is empty"))
	}

	return &model.Property{
		ID:  id,
		URL: url,
	}, nil
}

// InsertCaptureByPropertyID x
func (s *PropertyRedisStore) InsertCaptureByPropertyID(ctx context.Context, propID string, c *model.Capture) error {
	captureKey := propID + s.captureSuffix
	bs, err := json.Marshal(c)
	if err != nil {
		return errors.Wrap(perr.NewErrInternal(err), "could not marshal Capture to JSON")
	}

	if err := s.client.LPush(captureKey, bs).Err(); err != nil {
		return errors.Wrap(perr.NewErrInternal(err), "could not execute Redis command")
	}

	return nil
}

// InsertProperty x
func (s *PropertyRedisStore) InsertProperty(ctx context.Context, p *model.Property) error {
	currMaxID, err := s.client.Incr(s.idCounterKey).Result()
	if err != nil {
		return errors.Wrap(perr.NewErrInternal(err), "could not increment ID counter")
	}

	base16ID := fmt.Sprintf("%0"+s.idWidth+"x", currMaxID)
	if err := s.client.HSet(base16ID, map[string]interface{}{
		"id":  base16ID,
		"url": p.URL,
	}).Err(); err != nil {
		return errors.Wrap(perr.NewErrInternal(err), "could not set Property hashmap")
	}

	return nil
}
