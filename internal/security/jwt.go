package security

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const signingKey = "secret"

type UserClaims struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	jwt.StandardClaims
}

func generateToken(id uuid.UUID, email string) (string, error) {
	claims := UserClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}
