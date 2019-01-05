package helpers

import "net/http"

// ResponseFromStruct create response from struct
func ResponseFromStruct(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
