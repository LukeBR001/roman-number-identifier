package pkg

import (
	"identifier/pkg/model"
	"net/http"
)

func RomanIdentify(w http.ResponseWriter, r *http.Request) {
	var randomTextPayload model.TextPayload

	err := ReadPayloadRequest(r, &randomTextPayload)
	if err != nil {
		return
	}

	res, err := CountRomanNumber(randomTextPayload)
	if err != nil {
		return
	}

	WritePayloadResponse(w, http.StatusOK, res)
	return
}
