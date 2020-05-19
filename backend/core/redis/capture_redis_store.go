package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"zcrapr/core/model"
	"zcrapr/core/perr"
	"zcrapr/core/plog"

	goredis "github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

// CaptureRedisStore is a model.CaptureStore backed by Redis
type CaptureRedisStore struct {
	l      plog.Logger
	client *goredis.Client

	captureSuffix string
}

// NewCaptureRedisStore returns a new CaptureStore backed by Redis
func NewCaptureRedisStore(l plog.Logger, captureSuffix, host, password string, port uint) (*CaptureRedisStore, error) {
	if captureSuffix == "" {
		return nil, perr.NewErrInvalid("captureSuffix cannot be empty string")
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

	return &CaptureRedisStore{
		l:             l,
		client:        client,
		captureSuffix: captureSuffix,
	}, nil
}

// GetAllCapturesByPropertyID x
func (s *CaptureRedisStore) GetAllCapturesByPropertyID(ctx context.Context, propID string) ([]model.Capture, error) {
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
