package model

type TextPayload struct {
	Text string `json:"text"`
}

type DefinedReferenceSymbols struct {
	Combinations map[string]int
}
