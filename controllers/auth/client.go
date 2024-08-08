package auth

import (
	c "blops-me/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	config = oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	jwtSecret []byte
)
