package hash

import (
	"auth/entity"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func CreateSignature(p []byte) (string, error) {
	mac := hmac.New(sha256.New, []byte(os.Getenv("SECRET_KEY")))
	_, err := mac.Write(p)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func Verify(msg []byte, hash string) (bool, error) {
	sig, err := hex.DecodeString(hash)
	fmt.Println(sig)
	if err != nil {
		return false, err
	}
	log.Println(hash)
	mac := hmac.New(sha256.New, []byte(os.Getenv("SECRET_KEY")))
	mac.Write(msg)

	return hmac.Equal(sig, mac.Sum(nil)), nil
}

func Validate(msg []byte, hash string) (bool, error) {
	p, err := CreateSignature(msg)
	if err != nil {
		return false, err
	}
	return p == hash, nil
}

func CreateToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": user.Email,
			"id":    user.ID,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
