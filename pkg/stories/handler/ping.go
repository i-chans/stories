package handler

import "net/http"

func PingHandler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		writeResponse(resp, []byte("pong"), http.StatusOK)
	}
}
