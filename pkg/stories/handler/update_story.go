package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"stories/pkg/stories/contract"
	"stories/pkg/stories/service"
)

func UpdateStory(lgr *zap.Logger, svc *service.Service) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var updateReq contract.UpdateStoryRequest
		err := parseRequest(resp, req, lgr, &updateReq)
		if err != nil {
			return
		}

		updateResp, err := svc.StoriesService().Update(updateReq)
		if err != nil {
			handleError(resp, err, http.StatusInternalServerError, false, lgr)
		}

		respData, err := json.Marshal(updateResp)
		if err != nil {
			handleError(resp, err, http.StatusInternalServerError, true, lgr)
			return
		}

		writeResponse(resp, respData, http.StatusOK)
	}
}
