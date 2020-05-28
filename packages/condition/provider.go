package condition

type Condition struct {
	IHandler
}

func NewCondition() *Condition {
	rp := NewRepository()
	uc := NewUsecase(rp)
	h := NewHandler(uc)
	return &Condition{h}
}
