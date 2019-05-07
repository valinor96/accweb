package api

import (
	"github.com/assetto-corsa-web/accweb/server"
	"net/http"
	"strconv"
)

func StartInstanceHandler(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Id int `json:"id"`
	}{}

	if err := decodeJSON(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, nil)
		return
	}

	if err := server.StartServer(req.Id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, nil)
		return
	}

	writeResponse(w, nil)
}

func StopInstanceHandler(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Id int `json:"id"`
	}{}

	if err := decodeJSON(r, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, nil)
		return
	}

	if err := server.StopServer(req.Id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, nil)
		return
	}

	writeResponse(w, nil)
}

func GetInstanceLogsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, nil)
		return
	}

	logs, err := server.GetServerLogs(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, nil)
		return
	}

	resp := &struct {
		Logs string `json:"logs"`
	}{logs}
	writeResponse(w, resp)
}
