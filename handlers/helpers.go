package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type Map map[string]any

func writeJson(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to marshal data", "error", err)
		return
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		slog.Error("failed to write response", "error", err)
		return
	}
}

func readJson(w http.ResponseWriter, r *http.Request, dest any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("Content-Type of 'application/json' expected")
	}

	maxBytes := 1048576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dest)
	if err != nil {
		return err
	}

	return nil
}

func sendServerError(w http.ResponseWriter) {
	writeJson(w, http.StatusInternalServerError, Map{
		"message": "the server could not process your request",
	})
}
