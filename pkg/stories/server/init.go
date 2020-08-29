package server

import (
	"go.uber.org/zap"
	"log"
	"stories/cmd/config"
	"stories/pkg/stories/service"
	"stories/pkg/stories/store"
)

func getService(cfg config.Config, lgr *zap.Logger) *service.Service {
	db, err := store.NewDBHandler(cfg.GetDatabaseConfig(), lgr).GetDB()
	if err != nil {
		log.Fatal(err)
	}

	return service.NewService(
		service.NewStoriesService(
			cfg.GetStoryConfig(), lgr, store.NewStore(store.NewStoriesStore(db, lgr)),
		),
	)
}
