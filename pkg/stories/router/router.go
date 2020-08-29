package router

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"stories/pkg/stories/handler"
	"stories/pkg/stories/service"
)

const (
	pingPath = "/ping"

	addStoryPath    = "/add"
	searchStoryPath = "/search"
	updateStoryPath = "/update"
	deleteStoryPath = "/delete"
)

func NewRouter(lgr *zap.Logger, svc *service.Service) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc(pingPath, handler.PingHandler()).Methods(http.MethodGet)

	r.HandleFunc(addStoryPath, handler.AddStory(lgr, svc)).Methods(http.MethodPost)
	r.HandleFunc(searchStoryPath, handler.SearchStory(lgr, svc)).Methods(http.MethodGet)
	r.HandleFunc(updateStoryPath, handler.UpdateStory(lgr, svc)).Methods(http.MethodPatch)
	r.HandleFunc(deleteStoryPath, handler.DeleteStory(lgr, svc)).Methods(http.MethodDelete)

	return r
}
