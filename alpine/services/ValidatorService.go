package services

import (
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

type ValidatorService struct{
}

const jwtKey = "YOUR_JWT_KEY" // TODO : get key from Vault

func (service *ValidatorService) VerifyToken(r *http.Request, w http.ResponseWriter) bool {

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		http.Error(w, "Permission denied", 401)
    	return false
	}

	reqToken = strings.TrimSpace(splitToken[1])

    token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
        return []byte(jwtKey), nil
    })
    if err == nil && token.Valid {
        fmt.Println("valid token found")
        return true
    } else {
        http.Error(w, "Permission denied", 401)
        return false
    }
}