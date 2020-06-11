package session

type Session struct {
	IHandler
}

func NewSession() *Session {
	r := NewRepository()
	u := NewUsecase(r)
	h := NewHandler(u)
	return &Session{h}
}
