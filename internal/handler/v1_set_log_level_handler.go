package handler

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
)

func (h *Handler) SetLogLevelV1(w http.ResponseWriter, r *http.Request) {
	level := r.URL.Query().Get("level")
	parsedLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		h.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	zerolog.SetGlobalLevel(parsedLevel)

	fmt.Println("Set log level to", level)
}
