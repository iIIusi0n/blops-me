package auth

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"io"
	"log"
)

type GoogleUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getUserInfo(token *oauth2.Token) (*GoogleUser, error) {
	client := config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close response body: ", err)
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user GoogleUser
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
