package server

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"stories/cmd/config"
	"stories/pkg/stories/router"
	"syscall"
	"time"
)

type Server interface {
	Start()
}

type storiesServer struct {
	cfg config.Config
	lgr *zap.Logger
}

func (ss *storiesServer) Start() {
	svc := getService(ss.cfg, ss.lgr)
	server := newHttpServer(ss.cfg.GetServerConfig(), router.NewRouter(ss.lgr, svc))

	ss.lgr.Sugar().Info("listening on: ", ss.cfg.GetServerConfig().GetAddress())
	go func() { _ = server.ListenAndServe() }()

	waitForShutDown(ss.lgr, server)
}

func waitForShutDown(lgr *zap.Logger, server *http.Server) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	err := server.Shutdown(context.Background())
	if err != nil {
		lgr.Error(err.Error())
		return
	}

	lgr.Info("server shutdown successful")
}

func newHttpServer(cfg config.ServerConfig, router *mux.Router) *http.Server {
	return &http.Server{
		Addr:         cfg.GetAddress(),
		Handler:      router,
		ReadTimeout:  time.Second * time.Duration(cfg.GetReadTimeOut()),
		WriteTimeout: time.Second * time.Duration(cfg.GetWriteTimeOut()),
	}
}

func NewServer(cfg config.Config, lgr *zap.Logger) Server {
	return &storiesServer{
		cfg: cfg,
		lgr: lgr,
	}
}
