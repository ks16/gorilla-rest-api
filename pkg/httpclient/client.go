package httpclient

import (
	"context"
	"github.com/ks16/gorilla-rest-api/pkg/middleware"
	"net/http"
)

type Client struct {
	*http.Client

	ctx context.Context
}

func NewClient(ctx context.Context) *Client {
	return &Client{&http.Client{}, ctx}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	correlationId, ok := c.ctx.Value(middleware.ContextKey(middleware.CorrelationIdKey)).(string)
	if ok {
		req.Header.Set(middleware.CorrelationHeaderName, correlationId)
	}

	return c.Client.Do(req)
}
