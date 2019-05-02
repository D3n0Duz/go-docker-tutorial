packages models

import time

type AccountModel struct {
	AccountStateID int
	AccountTypeID int
	Closed time.Time
	Created time.Time
	Email string
	LastConnection time.Time
}