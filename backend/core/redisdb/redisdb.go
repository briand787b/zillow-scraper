package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"zcrapr/core/model"
	"zcrapr/core/plog"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

// Client is
type Client struct {
	l      plog.Logger
	client redis.Client
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
		return err
	}

	return &Client{
		l:      l,
		client: client,
	}, nil
}

// AddCapture adds a capture event to Redis
func (c *Client) AddCapture(ctx context.Context, url string, c *model.Capture) error {
	bs, err := json.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "could not marshal Capture")
	}

	if err := c.client.LPush(url, string(bs)).Err(); err != nil {
		return errors.Wrap(err, "could not execute Redis command")
	}

	return nil
}

func (c *Client) GetCaptures(ctx context.Context, a *Property) ([]Capture, error) {
	c.client.L
}
