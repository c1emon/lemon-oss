package manager

import (
	"fmt"
	"strconv"

	"github.com/c1emon/gcommon/httpx"
	"github.com/c1emon/lemon_oss/internal/manager/store"
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

		Type       store.OSSType `json:"type"`
		Tls        bool          `json:"tls"`
		Endpoint   string        `json:"endpoint"`
		AccessKey  string        `json:"access_key"`
		AccessId   string        `json:"access_id"`
		BucketName string        `json:"bucket_name"`
	}{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData("json parse err"))
		return
	}

	p := &store.OSSProvider{
		Name:       param.Name,
		Oid:        param.Oid,
		Type:       param.Type,
		Tls:        param.Tls,
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

func (h *Handlers) InitUploadhandler(c *gin.Context) {
	providerId := c.Param("provider_id")
	mulitPart, err := strconv.Atoi(c.Query("mulit_part"))
	if err != nil {
		c.JSON(200, httpx.NewResponse(200).WithData(fmt.Sprintf("mulit_part invaild: %s", err)))
		return
	}

	id := h.m.InitUpload(providerId, mulitPart > 1)

	ret := &struct {
		Id string `json:"id"`
	}{Id: id}
	c.JSON(200, httpx.ResponseOK().WithData(ret))
}

func (h *Handlers) UploadHandler(c *gin.Context) {
	reqId := c.Param("req_id")

	param := &struct {
		Object string `form:"object"`
		Name   string `form:"name"`
	}{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData(fmt.Sprintf("param parse error: %s", err)))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData(fmt.Sprintf("retrive form file error: %s", err)))
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData(fmt.Sprintf("read form file error: %s", err)))
		return
	}
	defer f.Close()

	err = h.m.Upload(reqId, param.Object, f, file.Size)
	if err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData(fmt.Sprintf("upload file error: %s", err)))
		return
	}

	c.JSON(200, httpx.ResponseOK())
}

func (h *Handlers) CompleteUploadHandler(c *gin.Context) {
	reqId := c.Param("req_id")

	req, err := h.m.CompleteUpload(reqId)
	if err != nil {
		c.JSON(200, httpx.NewResponse(1).WithData(fmt.Sprintf("complete upload error: %s", err)))
		return
	}

	c.JSON(200, httpx.ResponseOK().WithData(req))
}
