package prediction

import "time"

type UserPrediction struct {
	Prediction int
	ImagePath  string
	UserID     int
	UpdatedAt  time.Time
}
