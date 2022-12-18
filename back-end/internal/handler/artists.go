package handler

import (
	"groupie-tracker/back-end/pkg/errors"
	"groupie-tracker/back-end/pkg/webjson"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) Artist(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		h.logger.Err.Println("Only get method in h.Artist")
		webjson.JSONError(w, errors.WebFail(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/artists/")
	if idStr == "" {
		h.GetAllGroups(w, r)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Err.Printf("%s not found\n", r.URL.Path)
		webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	h.GetGroupById(w, r, id)
}

func (h *Handler) GetGroupById(w http.ResponseWriter, r *http.Request, id int) {
	group, err := h.service.Group.GetById(id)
	if err != nil {
		h.logger.Err.Printf("internal server error")
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if group.Id == 0 {
		h.logger.Err.Printf("not accebtable id - %d", id)
		webjson.JSONError(w, errors.WebFail(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	webjson.SendJSON(w, group)
}

func (h *Handler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	group, err := h.service.Group.GetAll()
	if err != nil {
		h.logger.Err.Printf("internal server error: %s", err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	webjson.SendJSON(w, group)
}
