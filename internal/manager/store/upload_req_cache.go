package store

import (
	"github.com/c1emon/gcommon/cachex"
	"github.com/c1emon/gcommon/cachex/memory"
)

func NewUploadRequestCacher() *UploadReqCacher {
	return &UploadReqCacher{
		cache: memory.NewMemoryCache(),
	}
}

type UploadReqCacher struct {
	cache cachex.Cacher
}

func (c *UploadReqCacher) Get(key string) *UploadRequest {
	if req, ok := c.cache.Get(key); ok {
		return req.(*UploadRequest)
	} else {
		return nil
	}
}

func (c *UploadReqCacher) Set(key string, req *UploadRequest) {
	c.cache.Set(key, req)
}

func (c *UploadReqCacher) Del(key string) {
	c.cache.Del(key)
}
