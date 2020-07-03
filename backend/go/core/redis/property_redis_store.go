package redis

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"zcrapr/core/model"
	"zcrapr/core/perr"
	"zcrapr/core/plog"

	"github.com/go-redis/redis/v7"
	goredis "github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

// PropertyRedisStore is a PropertyStore backed by Redis
type PropertyRedisStore struct {
	l      plog.Logger
	client *goredis.Client

	idCounterKey string
	idPrefix     string
	idWidth      string
	maxCursor    uint
}

// NewPropertyRedisStore returns a new PropertyStore backed by Redis
func NewPropertyRedisStore(l plog.Logger, idCounterKey, idPrefix, host, password string,
	idWidth, maxCursor, port uint) (*PropertyRedisStore, error) {

	if idCounterKey == "" {
		return nil, perr.NewErrInvalid("idCounterKey cannot be empty string")
	}

	if idPrefix == "" {
		return nil, perr.NewErrInvalid("idPrefix cannot be empty string")
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
		l:            l,
		client:       client,
		idCounterKey: idCounterKey,
		idPrefix:     idPrefix,
		idWidth:      strconv.Itoa(int(idWidth)),
	}, nil
}

// GetAllPropertyIDs x
func (s *PropertyRedisStore) GetAllPropertyIDs(ctx context.Context, take int) ([]string, error) {
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

// GetPropertiesByAddress x
func (s *PropertyRedisStore) GetPropertiesByAddress(ctx context.Context, address string) ([]model.Property, error) {
	var allKeys []string
	var nextKeys []string
	var err error
	var cursor uint64
	for {
		nextKeys, cursor, err = s.client.Scan(cursor, "*"+strings.ToLower(address)+"*", math.MaxInt32).Result()
		if err != nil {
			return nil, perr.NewErrInternal(errors.Wrap(err, "could not execute Redis command"))
		}

		allKeys = append(allKeys, nextKeys...)
		if cursor == 0 {
			break
		}
	}

	addrs := s.removeURLs(allKeys)
	addrs = s.removeIDs(addrs)

	var id string
	var prop *model.Property
	props := make([]model.Property, len(addrs))
	for i, addr := range addrs {
		id, err = s.GetPropertyIDByAddress(ctx, addr)
		if err != nil {
			return nil, errors.Wrap(err, "could not get propertyID by address")
		}

		prop, err = s.GetPropertyByID(ctx, id)
		if err != nil {
			return nil, errors.Wrap(err, "could not get property by ID")
		}

		props[i] = *prop
	}

	return props, nil
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
	if !ok || url == "" {
		return nil, perr.NewErrInternal(errors.New("returned url is empty"))
	}

	acreageStr, ok := fieldMap["acreage"]
	if !ok || acreageStr == "" {
		return nil, perr.NewErrInternal(errors.New("returned acreage is empty"))
	}

	acreage, err := strconv.Atoi(acreageStr)
	if err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not convert acreage to int")
	}

	address, ok := fieldMap["address"]
	if !ok || address == "" {
		return nil, perr.NewErrInternal(errors.New("returned address is empty"))
	}

	return &model.Property{
		ID:      id,
		URL:     url,
		Acreage: acreage,
		Address: address,
	}, nil
}

// GetPropertyIDByAddress x
func (s *PropertyRedisStore) GetPropertyIDByAddress(ctx context.Context, address string) (string, error) {
	id, err := s.client.Get(address).Result()
	if err != nil {
		if err == redis.Nil {
			return "", perr.NewErrNotFound(errors.New("address does not exist in database"))
		}

		return "", errors.Wrap(perr.NewErrInternal(err), "could not get id by address")
	}

	return id, nil
}

// GetPropertyIDByURL x
func (s *PropertyRedisStore) GetPropertyIDByURL(ctx context.Context, url string) (string, error) {
	id, err := s.client.Get(url).Result()
	if err != nil {
		if err == redis.Nil {
			return "", perr.NewErrNotFound(errors.New("url does not exist in database"))
		}

		return "", errors.Wrap(perr.NewErrInternal(err), "could not get id by url")
	}

	return id, nil
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
		"id":      base16ID,
		"url":     p.URL,
		"acreage": p.Acreage,
		"address": strings.ToLower(p.Address),
	})
	txPipe.Set(p.URL, base16ID, 0)
	txPipe.Set(strings.ToLower(p.Address), base16ID, 0)
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

////// HELPER METHODS //////

func (s *PropertyRedisStore) removeIDs(keyList []string) []string {
	var stripped = make([]string, 0, len(keyList))
	for _, key := range keyList {
		if !strings.Contains(key, "id:") {
			stripped = append(stripped, key)
		}
	}

	return stripped
}

func (s *PropertyRedisStore) removeURLs(keyList []string) []string {
	var stripped = make([]string, 0, len(keyList))
	for _, key := range keyList {
		if !strings.Contains(key, "http") {
			stripped = append(stripped, key)
		}
	}

	return stripped
}
