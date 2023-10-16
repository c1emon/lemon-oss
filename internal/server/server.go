package server

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/c1emon/gcommon/logx"
	"github.com/c1emon/lemon_oss/internal/setting"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func Initialize(cfg *setting.Config) (*Server, error) {
	srv, _ := ProvideHttpServer(cfg)
	return New(cfg, srv)
}

// from https://github.com/grafana/grafana/blob/4cc72a22ad03132295ab3428ed9877ba2cb42eb2/pkg/server/server.go
func New(cfg *setting.Config, httpSrv *HttpServer) (*Server, error) {
	s, err := newServer(cfg, httpSrv)
	if err != nil {
		return nil, err
	}

	if err := s.Init(); err != nil {
		return nil, err
	}

	return s, nil
}

func newServer(cfg *setting.Config, httpSrv *HttpServer) (*Server, error) {
	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)

	s := &Server{
		context:          childCtx,
		childRoutines:    childRoutines,
		shutdownFn:       shutdownFn,
		shutdownFinished: make(chan any),
		log:              logx.GetLogger(),
		cfg:              cfg,
		httpSrv:          httpSrv,
	}

	return s, nil
}

type Server struct {
	context          context.Context
	shutdownFn       context.CancelFunc
	childRoutines    *errgroup.Group
	log              *logrus.Logger
	cfg              *setting.Config
	shutdownOnce     sync.Once
	shutdownFinished chan any
	isInitialized    bool
	mtx              sync.Mutex

	// pidFile     string
	// version     string
	// commit      string
	// buildBranch string

	httpSrv *HttpServer

	// backgroundServices []registry.BackgroundService
}

func (s *Server) Init() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.isInitialized {
		return nil
	}
	s.isInitialized = true

	return nil
}

func (s *Server) Run() error {
	defer close(s.shutdownFinished)

	if err := s.Init(); err != nil {
		return err
	}

	s.childRoutines.Go(func() error {
		select {
		case <-s.context.Done():
			return s.context.Err()
		default:
		}

		// start service
		s.log.Debugf("Starting background service: %s", "http server")
		err := s.httpSrv.Run(s.context)
		// Do not return context.Canceled error since errgroup.Group only
		// returns the first error to the caller - thus we can miss a more
		// interesting error.
		if err != nil && !errors.Is(err, context.Canceled) {
			s.log.Errorf("Stopped background service: %s for %s", "http server", err)
			return fmt.Errorf("%s run error: %w", "http server", err)
		}
		s.log.Debugf("Stopped background service %s for %s", "http server", err)
		return nil

	})

	return s.childRoutines.Wait()
}

// Shutdown initiates Grafana graceful shutdown. This shuts down all
// running background services. Since Run blocks Shutdown supposed to
// be run from a separate goroutine.
func (s *Server) Shutdown(ctx context.Context, reason string) error {
	var err error
	s.shutdownOnce.Do(func() {
		s.log.Infof("Shutdown started: %s", reason)
		// Call cancel func to stop background services.
		s.shutdownFn()
		// Wait for server to shut down
		select {
		case <-s.shutdownFinished:
			s.log.Debug("Finished waiting for server to shut down")
		case <-ctx.Done():
			s.log.Warn("Timed out while waiting for server to shut down")
			err = fmt.Errorf("timeout waiting for shutdown")
		}
	})

	return err
}
