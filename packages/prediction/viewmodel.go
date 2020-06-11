package prediction

import "time"

type PredictionVM struct {
	Prediction int
	Label      string
	ImagePath  string
	UserID     int
	UpdatedAt  time.Time
}
