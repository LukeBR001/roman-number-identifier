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

	identifiedCombinations, err := IdentifyCombinations(randomTextPayload)
	if err != nil {
		WritePayloadResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res := IdentifyBiggerNumber(identifiedCombinations)

	WritePayloadResponse(w, http.StatusOK, res)
	return
}
