package main

import "net/http"

func sendError(w http.ResponseWriter, message string) {
	output := "{\"error\":\"" + message + "\"}"
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(output))
}
