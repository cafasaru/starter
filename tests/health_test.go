// +build e2e

package tests

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("missionimpossible"))
	if err != nil {
		fmt.Println(err)

	}

	return tokenString
}

func TestHealthCheck(t *testing.T) {
	t.Run("health check", func(t *testing.T) {
		client := resty.New()

		resp, err := client.R().Get("http://localhost:8080/alive")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

}
