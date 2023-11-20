// auth.go

package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("OTg5o6vMrVQUouyWv6fOcrFAgBLjKlnuzv7q0UZtVRw=")

// Claims representa la estructura de los claims en el token
type Claims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}

// GenerateAuthToken genera un token de autenticación (cambiado a mayúscula inicial)
func GenerateAuthToken(userID int) (string, error) {
	// Crea un conjunto de claims
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Tiempo de expiración del token (1 hora)
		},
	}

	// Crea un token firmado con los claims y la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
