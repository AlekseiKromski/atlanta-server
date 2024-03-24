package guard

import (
	"alekseikromski.com/atlanta/modules/storage"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type Guard struct {
	secret       []byte
	permissions  map[string][]*storage.Endpoint
	store        storage.Storage
	cookieDomain string
}

func NewGuard(secret []byte, store storage.Storage, cookieDomain string) *Guard {
	permissions, err := store.GetPermissions()
	if err != nil {
		log.Printf("cannot get permissions: %v", err)
	}

	return &Guard{
		secret:       secret,
		store:        store,
		permissions:  permissions,
		cookieDomain: cookieDomain,
	}
}

func (g *Guard) Auth(c *gin.Context) {
	defer c.Request.Body.Close()

	ar := &authRequest{}
	if err := json.NewDecoder(c.Request.Body).Decode(&ar); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	user, err := g.store.GetUserByUsername(ar.Username)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ar.Password)); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["authorized"] = true
	claims["id"] = user.Id

	tokenString, err := token.SignedString(g.secret)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	if ar.Type == "cookie" {
		c.SetCookie("token", tokenString, 3600, "/", g.cookieDomain, true, true)
	} else {
		c.JSON(http.StatusOK, struct {
			Token string `json:"token"`
		}{
			Token: tokenString,
		})
	}

	return
}

func (g *Guard) Check(c *gin.Context) {
	req := c.Request
	tokenRequest := ""
	t, err := c.Cookie("token")
	if err != nil {
		log.Printf("[JWTGUARD] there is problem to get token from cookies")
	}

	if t != "" {
		tokenRequest = t
	} else {
		tokenRequest = req.Header.Get("Authorization")
		if len(tokenRequest) != 0 {
			tokenRequest = tokenRequest[7:len(tokenRequest)]
		}
	}

	userID := ""
	if tokenRequest == "" || len(tokenRequest) < 10 {
		log.Printf("[JWTGUARD] there is not token in request: %s", req.URL.String())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenRequest, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong sign method")
		}
		claims := token.Claims.(jwt.MapClaims)
		if claims["id"] == nil {
			return nil, fmt.Errorf("wrong format of JWT")
		}

		userid, ok := claims["id"].(string)
		if !ok {
			return nil, fmt.Errorf("wrong format of userid")
		}

		userID = userid

		return g.secret, nil
	})

	if err != nil {
		log.Printf("[JWTGUARD] token verify failed: %v", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := g.store.GetUserById(userID)
	if err != nil {
		log.Printf("[JWTGUARD] cannot find user by id %s: %v", userID, err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	denied := true
	for _, path := range g.permissions[user.Role] {
		if req.URL.Path == path.Urn {
			denied = false
		}
	}

	if denied {
		log.Printf("[JWTGUARD] access denied by permission restrictions. User role %s / rights: %s", user.Role, g.permissions[user.Role])
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	// Set for at-socket-server
	c.Set("uid", userID)
	c.Next()
}
