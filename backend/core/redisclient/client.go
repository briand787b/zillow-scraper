package redisclient

import (
	"context"
	"encoding/json"
	"fmt"

	"zcrapr/core/model"
	"zcrapr/core/perr"
	"zcrapr/core/plog"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

// Client is
type Client struct {
	l      plog.Logger
	client *redis.Client
}

// NewClient returns a new client
func NewClient(l plog.Logger, host, password string, port uint) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", host, port),
		Password: password,
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		l.Error(nil, pong, "error", err)
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not ping Redis")
	}

	return &Client{
		l:      l,
		client: client,
	}, nil
}

// AddCapture adds a capture event to Redis
func (c *Client) AddCapture(ctx context.Context, url string, cap *model.Capture) error {
	bs, err := json.Marshal(cap)
	if err != nil {
		return errors.Wrap(perr.NewErrInternal(err), "could not marshal Capture to JSON")
	}

	if err := c.client.LPush(url, string(bs)).Err(); err != nil {
		return errors.Wrap(perr.NewErrInternal(err), "could not execute Redis command")
	}

	return nil
}

// GetLatestCapture retrieves the latest capture
func (c *Client) GetLatestCapture(ctx context.Context, url string) (*model.Capture, error) {
	caps, err := c.client.LRange(url, 0, 0).Result()
	if err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not execute Redis command")
	}

	if len(caps) < 1 {
		return nil, perr.NewErrNotFound(errors.Errorf("no captures for url %s", url))
	}

	var cap model.Capture
	if err := json.Unmarshal([]byte(caps[0]), &cap); err != nil {
		return nil, errors.Wrap(perr.NewErrInternal(err), "could not unmarshal to Redis output to Capture")
	}

	return &cap, nil
}
