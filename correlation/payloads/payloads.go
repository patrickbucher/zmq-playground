package payloads

type RequestPayload struct {
	A             int    `json:"a"`
	B             int    `json:"b"`
	CorrelationID string `json:"correlationId"`
}

type ResponsePayload struct {
	C             int    `json:"c"`
	CorrelationID string `json:"correlationId"`
}
