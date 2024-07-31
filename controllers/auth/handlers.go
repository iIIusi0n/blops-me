package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func LoginHandler(c *gin.Context) {
	url := config.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackHandler(c *gin.Context) {
	token, err := config.Exchange(c, c.Query("code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exchange token"})
		return
	}

	user, err := getUserInfo(token)
	if err != nil {
		log.Println("Failed to get user info: ", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := jwtToken.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", tokenString, 3600*24, "/", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/s")
}

func VerifyHandler(c *gin.Context) {
	ok := c.MustGet("authorized").(bool)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"authorized": false})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"authorized": true})
		return
	}
}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
