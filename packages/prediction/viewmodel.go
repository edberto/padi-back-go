package prediction

import "time"

type PredictionVM struct {
	Prediction int
	ImagePath  string
	UserID     int
	UpdatedAt  time.Time
}
