package register

type Register struct {
	IHandler
}

func NewRegister() *Register {
	r := NewRepo()
	u := NewUsecase(r)
	h := NewHandler(u)
	return &Register{h}
}
