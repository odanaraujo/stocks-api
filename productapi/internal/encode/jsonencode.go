package encode

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, obj interface{}, status int) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "JSON")
	w.WriteHeader(status)
	w.Write(bytes)
}
