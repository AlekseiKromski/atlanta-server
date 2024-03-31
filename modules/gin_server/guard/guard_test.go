package guard

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestGuard_Auth(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("Jusebox200"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash))
}
