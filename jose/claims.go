package jose

import (
	"time"

	"github.com/pkg/errors"

	"github.com/mj23978/chat-backend-x/utils"
)

// Claims represents a JSON Web Token's standard claims.
type Claims struct {
	// Audience identifies the recipients that the JWT is intended for.
	Audience []string `json:"aud"`

	// Issuer identifies the principal that issued the JWT.
	Issuer string `json:"iss"`

	// Subject identifies the principal that is the subject of the JWT.
	Subject string `json:"sub"`

	// ExpiresAt identifies the expiration time on or after which the JWT most not be accepted for processing.
	ExpiresAt time.Time `json:"exp"`

	// IssuedAt identifies the time at which the JWT was issued.
	IssuedAt time.Time `json:"iat"`

	// NotBefore identifies the time before which the JWT must not be accepted for processing.
	NotBefore time.Time `json:"nbf"`

	// JTI provides a unique identifier for the JWT.
	JTI string `json:"jti"`
}

// ParseMapStringInterfaceClaims converts map[string]interface{} to *Claims.
func ParseMapStringInterfaceClaims(claims map[string]interface{}) *Claims {
	c := make(map[interface{}]interface{})
	for k, v := range claims {
		c[k] = v
	}
	return ParseMapInterfaceInterfaceClaims(c)
}

// ParseMapInterfaceInterfaceClaims converts map[interface{}]interface{} to *Claims.
func ParseMapInterfaceInterfaceClaims(claims map[interface{}]interface{}) *Claims {
	result := &Claims{
		Issuer:  utils.GetStringDefault(claims, "iss", ""),
		Subject: utils.GetStringDefault(claims, "sub", ""),
		JTI:     utils.GetStringDefault(claims, "jti", ""),
	}

	if aud, err := utils.GetString(claims, "aud"); err == nil {
		result.Audience = []string{aud}
	} else if errors.Cause(err) == utils.ErrKeyCanNotBeTypeAsserted {
		if aud, err := utils.GetStringSlice(claims, "aud"); err == nil {
			result.Audience = aud
		} else {
			result.Audience = []string{}
		}
	} else {
		result.Audience = []string{}
	}

	if exp, err := utils.GetTime(claims, "exp"); err == nil {
		result.ExpiresAt = exp
	}

	if iat, err := utils.GetTime(claims, "iat"); err == nil {
		result.IssuedAt = iat
	}

	if nbf, err := utils.GetTime(claims, "nbf"); err == nil {
		result.NotBefore = nbf
	}

	return result
}
