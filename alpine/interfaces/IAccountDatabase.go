package interfaces

import (
	firestore "cloud.google.com/go/firestore"
)

type IAccountDatabase interface {
	GetClientConnection() firestore.Client
}
