package auth

import (
	"crypto/rand"
	"log"
)

func init() {
	secret := make([]byte, 32)
	_, err := rand.Read(secret)
	if err != nil {
		log.Fatalln("Failed to generate JWT secret: ", err)
	}

	jwtSecret = secret

	log.Println("JWT secret generated")
}
