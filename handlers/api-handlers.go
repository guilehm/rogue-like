package handlers

import (
	"encoding/json"
	"net/http"
	"rogue-like/models"
)

func SpriteList(h *models.Hub, w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(h.PlayerSprites)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}
