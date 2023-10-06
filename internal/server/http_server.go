package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/c1emon/lemon_oss/internal/setting"
	"github.com/c1emon/lemon_oss/pkg/ginx"
	"github.com/c1emon/lemon_oss/pkg/logx"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ProvideHttpServer(cfg *setting.Config) (*HttpServer, error) {
	return &HttpServer{
		server: ginx.GetGinEng(),
		port:   cfg.Http.Port,
		log:    logx.GetLogger(),
	}, nil
}

type HttpServer struct {
	log    *logrus.Logger
	server *gin.Engine
	port   int
}

func (s *HttpServer) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(1)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.server,
	}

	// handle http shutdown on server context done
	go func() {
		defer wg.Done()
		<-ctx.Done()

		// shutdown server here
		srv_ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFn()
		if err := srv.Shutdown(srv_ctx); err != nil {
			s.log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		<-srv_ctx.Done()
		s.log.Println("timeout of 5 seconds.")

	}()

	RegRouter()

	// start server here
	if err := srv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			s.log.Debug("server was shutdown gracefully")
			return nil
		}
		return err
	}

	wg.Wait()
	return nil
}
