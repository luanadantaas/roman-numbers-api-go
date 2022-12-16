package entity

type Number struct {
	Number string `json:"number"`
	Value  int    `json:"value"`
}

type ApiResponse struct {
	Text string `json:"text"`
}
