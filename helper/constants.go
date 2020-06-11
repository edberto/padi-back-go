package helper

import (
	"errors"
	"time"
)

var (
	ErrUserExisted  = errors.New("User already existed!")
	ErrUserNotFound = errors.New("User not found!")
	ErrTokenExpired = errors.New("Token Expired!")

	LabelItoa = map[int]string{
		1: "Sakit (Blas/Busuk Leher/Patah Leher/Tekek/Kecekik)",
		2: "Sehat",
		3: "Sakit (Hispa)",
		4: "Sakit (Bercak Coklat)",
	}

	WIB = time.FixedZone("UTC+7", 7*60*60)
)
