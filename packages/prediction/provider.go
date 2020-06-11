package prediction

type Prediction struct {
	IHandler
}

func NewPrediction() *Prediction {
	r := NewRepository()
	u := NewUsecase(r)
	h := NewHandler(u)
	return &Prediction{h}
}
