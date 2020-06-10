package session

type User struct {
	ID       int
	Username string
	Password string
}

type Token struct {
	UUID      string `bson"uuid"`
	UserID    string `bson"user_id"`
	ExpiredAt string `bson"expired_at"`
}
