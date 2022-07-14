package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
