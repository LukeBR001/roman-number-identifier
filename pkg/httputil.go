package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadPayloadRequest(r *http.Request, payload interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &payload)

	return nil
}

func WritePayloadResponse(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(&payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}
