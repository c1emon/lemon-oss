package manager

import (
	"fmt"

	"github.com/c1emon/lemon_oss/pkg/httpx"
	"github.com/gin-gonic/gin"
)

func NewHandler() *Handlers {
	return &Handlers{NewManager()}
}

type Handlers struct {
	m *Manager
}

func (h *Handlers) CreateHandler(c *gin.Context) {
	param := &struct {
		Name string `json:"name"`
		Oid  string `json:"oid"`

		Type       OSSType `json:"type"`
		Endpoint   string  `json:"endpoint"`
		AccessKey  string  `json:"access_key"`
		AccessId   string  `json:"access_id"`
		BucketName string  `json:"bucket_name"`
	}{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData("json parse err"))
		return
	}

	p := &OSSProvider{
		Name:       param.Name,
		Oid:        param.Oid,
		Type:       param.Type,
		Endpoint:   param.Endpoint,
		AccessId:   param.AccessId,
		AccessKey:  param.AccessKey,
		BucketName: param.BucketName,
	}

	err := h.m.Create(p)
	if err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData(fmt.Sprintf("unable create oss provider: %s", err)))
		return
	}

	c.JSON(200, httpx.ResponseOK().WithData(p))
}

func (h *Handlers) STSHandler(c *gin.Context) {

	param := &struct {
		Id         string `json:"id"`
		ObjectName string `json:"obj"`
	}{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData("json parse err"))
		return
	}

	sts, err := h.m.GenSTS(param.Id, param.ObjectName)
	if err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData(fmt.Sprintf("STS gen err %s", err)))
		return
	}

	resp := &struct {
		STS string `json:"sts"`
	}{sts}

	c.JSON(200, httpx.ResponseOK().WithData(resp))
}
