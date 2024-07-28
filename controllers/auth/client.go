package auth

import (
	c "blops-me/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	config = oauth2.Config{
		ClientID:     c.CLIENT_ID,
		ClientSecret: c.CLIENT_SECRET,
		RedirectURL:  c.REDIRECT_URL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	jwtSecret []byte
)
