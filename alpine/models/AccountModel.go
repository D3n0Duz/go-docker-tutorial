package models

import (
	"time"
)

type AccountModel struct {
	Email              string
	Token              int
	Closed             time.Time
	Created            time.Time
	LastConnection     time.Time
	NumberOfConnection int
	Status             int
	Groups             int
}
