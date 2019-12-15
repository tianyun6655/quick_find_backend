package http


type SendMessageRequest struct {
	From int32  `json:"from"`
	To int32   `json:"to"`
	Message string `json:"message"`
}


type TranslateRequest struct {
	Text string `json:"text"`
}