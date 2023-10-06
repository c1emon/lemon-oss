package server

import (
	"github.com/c1emon/lemon_oss/internal/manager"
	"github.com/c1emon/lemon_oss/pkg/ginx"
)

func RegRouter() {
	h := manager.NewHandler()
	eng := ginx.GetGinEng()
	g1 := eng.Group("/api/v1/oss")
	g1.POST("/provider", h.CreateHandler)
	g1.POST("/sts", h.STSHandler)
}
