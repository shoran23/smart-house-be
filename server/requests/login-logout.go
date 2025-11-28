package requests

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Logout struct {
	Username string `json:"username"`
}
