package interfaces

import (
	auth "firebase.google.com/go/auth"
)

type IAccountDatabase interface {
	GetClientConnection() *auth.Client
}
