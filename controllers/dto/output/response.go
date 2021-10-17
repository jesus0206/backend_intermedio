package output

// A Response exported
type Response struct {
	Error   string `json:"error"`
	Mensaje string `json:"mensaje"`
}

// ResponseError response
type ResponseError struct {
	Error string `json:"error"`
}

// ResponseEmpty response
type ResponseEmpty struct {
	Data string `json:"data"`
}
