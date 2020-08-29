package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"stories/pkg/stories/contract"
	"stories/pkg/stories/service"
)

func DeleteStory(lgr *zap.Logger, svc *service.Service) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var deleteReq contract.DeleteStoryRequest
		err := parseRequest(resp, req, lgr, &deleteReq)
		if err != nil {
			return
		}

		deleteResp, err := svc.StoriesService().Delete(deleteReq)
		if err != nil {
			handleError(resp, err, http.StatusInternalServerError, false, lgr)
		}

		respData, err := json.Marshal(deleteResp)
		if err != nil {
			handleError(resp, err, http.StatusInternalServerError, true, lgr)
			return
		}

		writeResponse(resp, respData, http.StatusOK)
	}

}
