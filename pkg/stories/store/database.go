package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"stories/cmd/config"
	"time"
)

const databaseDriver = "postgres"

type DBHandler interface {
	GetDB() (*sql.DB, error)
}

type defaultDBHandler struct {
	lgr *zap.Logger
	cfg config.DatabaseConfig
}

func (ddh *defaultDBHandler) GetDB() (*sql.DB, error) {
	db, err := sql.Open(databaseDriver, ddh.cfg.GetDSN())
	if err != nil {
		ddh.lgr.Error(err.Error())
		return nil, err
	}

	db.SetMaxIdleConns(ddh.cfg.GetMaxIdleConnections())
	db.SetMaxOpenConns(ddh.cfg.GetMaxOpenConnections())
	db.SetConnMaxLifetime(time.Minute * time.Duration(ddh.cfg.GetConnectionMaxLifeTime()))

	err = db.Ping()
	if err != nil {
		ddh.lgr.Error(err.Error())
		return nil, err
	}

	return db, nil
}

func NewDBHandler(cfg config.DatabaseConfig, lgr *zap.Logger) DBHandler {
	return &defaultDBHandler{
		cfg: cfg,
		lgr: lgr,
	}
}
