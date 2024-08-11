package auth

import (
	"log"

	c "blops-me/config"
)

func init() {
	jwtSecret = []byte(c.SessionSecret)

	log.Printf("JWT secret set to: %v\n", string(jwtSecret))
}
