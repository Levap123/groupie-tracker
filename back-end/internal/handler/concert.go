package handler

import (
	"groupie-tracker/back-end/pkg/errors"
	"groupie-tracker/back-end/pkg/webjson"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) Concert(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		h.logger.Err.Println("Only get method in h.Concert")
		webjson.JSONError(w, errors.WebFail(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/concerts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Err.Printf("%s not found\n", r.URL.Path)
		webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	concert, err := h.service.Concert.GetByGroupId(id)
	if err != nil {
		h.logger.Err.Println("internal server error", id)
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if concert.Id == 0 {
		h.logger.Err.Printf("not accebtable id - %d", id)
		webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	webjson.SendJSON(w, concert)
}
