package smart_switch

type WebSocketRequest struct {
	Id     string `json:"id"`
	Src    string `json:"src"`
	Method string `json:"method"`
}
