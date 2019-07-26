package models

import (
	"time"
)

type AccountModel struct {
	AccountId string
	AccountStateId string
	AccountTypeId string
	Closed time.Time
	Created time.Time
	Email string
	LastConnection time.Time
}