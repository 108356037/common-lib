package middleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// type contextKey string

// const (
// 	uuidKey contextKey = "uuid"
// )

type AccessTokenPayload struct {
	Exp        float64 `json:"exp"`
	Sub        string  `json:"sub"`
	Event_id   string  `json:"event_id"`
	Auth_time  float64 `json:"auth_time"`
	Iss        string  `json:"iss"`
	Iat        float64 `json:"iat"`
	Jti        string  `json:"jti"`
	Client_id  string  `json:"client_id"`
	Username   string  `json:"username"`
	Origin_jti string  `json:"origin_jti"`
	Token_use  string  `json:"token_use"`
	Scope      string  `json:"scope"`
}

func RetriveDecodedJwt(b64str string) (*AccessTokenPayload, error) {

	if !strings.HasSuffix(b64str, "==") {
		b64str = b64str + "=="
	}

	dec, err := base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		return nil, err
	}

	var payload AccessTokenPayload
	err = json.Unmarshal(dec, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtPayload := c.Request.Header["Encoded-Jwt"]
		payload, err := RetriveDecodedJwt(jwtPayload[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"status": "failed",
				"info":   err.Error(),
			})
			c.Abort()
			return
		}

		//ctx := context.WithValue(c.Request.Context(), uuidKey, payload.Sub)
		//c.Request = c.Request.WithContext(ctx)

		c.Set("uuid", payload.Sub)
		c.Next()
	}
}
