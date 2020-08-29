package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"stories/pkg/stories/contract"
	"stories/pkg/stories/service"
)

func AddStory(lgr *zap.Logger, svc *service.Service) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var addReq contract.AddStoryRequest
		err := parseRequest(resp, req, lgr, &addReq)
		if err != nil {
			return
		}

		addResp, err := svc.StoriesService().Add(addReq)
		if err != nil {
			handleError(resp, err, http.StatusInternalServerError, false, lgr)
			return
		}

		respData, err := json.Marshal(addResp)
		if err != nil {
			handleError(resp, err, http.StatusInternalServerError, true, lgr)
			return
		}

		writeResponse(resp, respData, http.StatusOK)
	}
}
