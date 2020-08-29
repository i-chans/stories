package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func writeResponse(resp http.ResponseWriter, response []byte, statusCode int) {
	resp.WriteHeader(statusCode)
	_, _ = resp.Write(response)
}

func handleError(resp http.ResponseWriter, err error, statusCode int, log bool, lgr *zap.Logger) {
	if log {
		lgr.Error(err.Error())
	}

	writeResponse(resp, []byte(err.Error()), statusCode)
}

func parseRequest(resp http.ResponseWriter, req *http.Request, lgr *zap.Logger, reqStruct interface{}) error {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleError(resp, err, http.StatusBadRequest, true, lgr)
		return err
	}

	err = json.Unmarshal(data, reqStruct)
	if err != nil {
		handleError(resp, err, http.StatusBadRequest, true, lgr)
		return err
	}

	return nil
}
