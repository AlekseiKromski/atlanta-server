package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

func (v *V1) Auth(c *gin.Context) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * (time.Hour * 24)).Unix()
	claims["authorized"] = true
	claims["id"] = "test-user-id"

	tokenString, err := token.SignedString(v.secret)
	if err != nil {
		log.Fatalf("cannot create token: %v", err)
	}

	c.JSON(http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	})
}
