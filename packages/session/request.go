package session

type LoginR struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshR struct {
	Token string `json:"refresh_token"`
}
