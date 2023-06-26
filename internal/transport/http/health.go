package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {

	resp, err := h.Service.Health(r.Context())
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}

}
