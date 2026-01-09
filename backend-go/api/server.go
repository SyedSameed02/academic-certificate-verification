package api

import (
	"net/http"
	"strconv"

	"backend-go/config"
)

func StartServer(cfg *config.Config) error {
	addr := ":" + strconv.Itoa(cfg.Server.Port)
	return http.ListenAndServe(addr, NewRouter())
}
