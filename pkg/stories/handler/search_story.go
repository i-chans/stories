package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"stories/pkg/stories/contract"
	"stories/pkg/stories/service"
)

func SearchStory(lgr *zap.Logger, svc *service.Service) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var searchReq contract.SearchStoryRequest
		err := parseRequest(resp, req, lgr, &searchReq)
		if err != nil {
			return
		}

		searchResp, err := svc.StoriesService().Search(searchReq)
		if err != nil {
			handleError(resp, err, http.StatusBadRequest, false, lgr)
			return
		}

		respData, err := json.Marshal(searchResp)
		if err != nil {
			handleError(resp, err, http.StatusBadRequest, true, lgr)
			return
		}

		writeResponse(resp, respData, http.StatusOK)
	}
}
