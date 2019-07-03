package interfaces

import (
	"net/http"
)

type IValidatorService interface {
	VerifyToken(r *http.Request, w http.ResponseWriter) bool 
}
