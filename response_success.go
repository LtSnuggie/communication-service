package communication

import (
	"encoding/json"
	"net/http"
)

func ResponseSuccessMarshalJSON(w http.ResponseWriter, data interface{}) {
	ResponseSuccess(w)
	b, err := json.Marshal(data)
	if err != nil {
		ResponseErrorMarhsallingJSON(w)
		return
	}
	w.Write(b)
}

func ResponseSuccess(w http.ResponseWriter) {
	setHeaderForJSON(w)
	w.WriteHeader(http.StatusOK)
}
