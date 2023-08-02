package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func asApiResponse(jsonContent string, err error) []byte {
	if err != nil {
		return []byte(fmt.Sprintf(`{"content":null,"error": "%s"}`, err.Error()))
	}

	return []byte(fmt.Sprintf(`{"content":%s,"error": null}`, jsonContent))
}

func intURlParam(r *http.Request, key string) (int, error) {
	s := chi.URLParam(r, key)
	return strconv.Atoi(s)
}

func sendData(w http.ResponseWriter, data any) {
	raw, err := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	w.Write(asApiResponse(string(raw), err))
}

func sendErr(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.Write(asApiResponse("", err))
}

func send(w http.ResponseWriter, data any, err error) {
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		sendErr(w, err)
	} else {
		sendData(w, data)
	}
}
