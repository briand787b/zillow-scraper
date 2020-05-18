package redis

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
func NewPropertyRedisStore(l plog.Logger, idCounterKey, idPrefix, captureSuffix, host, password string,
	idWidth, port uint) (*PropertyRedisStore, error) {

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
	var cursor uint64
	count := int64(take)
	for {
		nextKeys, cursor, err = s.client.Scan(cursor, s.idPrefix+":*", count).Result()
		if err != nil {
			return nil, perr.NewErrInternal(errors.Wrap(err, "could not execute Redis command"))
		}

		allKeys = append(allKeys, nextKeys...)
		if cursor == 0 {
			break
		}

		if len(allKeys) >= take {
			break
		}
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

// GetPropertyByAddress x
func (s *PropertyRedisStore) GetPropertyByAddress(ctx context.Context, address string) (*model.Property, error) {
	id, err := s.client.Get(address).Result()
	if err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not get id by address")
	}

	return s.GetPropertyByID(ctx, id)
}

// GetPropertyByID x
func (s *PropertyRedisStore) GetPropertyByID(ctx context.Context, id string) (*model.Property, error) {
	fieldMap, err := s.client.HGetAll(id).Result()
	if err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not get all hash keys")
	}

	if len(fieldMap) < 1 {
		return nil, perr.NewErrNotFound(errors.Errorf("could not find Property with ID %s", id))
	}

	s.l.Info(ctx, "returned redis hash map", "hashmap", fieldMap)

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

// GetPropertyByURL x
func (s *PropertyRedisStore) GetPropertyByURL(ctx context.Context, url string) (*model.Property, error) {
	id, err := s.client.Get(url).Result()
	if err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not get id by url")
	}

	return s.GetPropertyByID(ctx, id)
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

	base16ID := fmt.Sprintf("%s:%0"+s.idWidth+"x", s.idPrefix, currMaxID)

	txPipe := s.client.TxPipeline()
	defer txPipe.Close()
	txPipe.HSet(base16ID, map[string]interface{}{
		"id":  base16ID,
		"url": p.URL,
	})
	txPipe.Set(p.URL, base16ID, 0)
	txPipe.Set(p.Address, base16ID, 0)
	if _, err := txPipe.ExecContext(ctx); err != nil {
		return errors.Wrap(perr.NewErrInternal(err), "could not execute transaction")
	}

	p.ID = base16ID

	return nil
}

// UpdateProperty x
func (s *PropertyRedisStore) UpdateProperty(ctx context.Context, p *model.Property) error {
	if err := s.client.HSet(p.ID, map[string]interface{}{
		"url": p.URL,
	}).Err(); err != nil {
		return errors.Wrap(perr.NewErrInternal(err), "could not execute redis command")
	}

	return nil
}
