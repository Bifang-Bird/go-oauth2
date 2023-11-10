package models

import (
	"fmt"
	"strings"

	oauth2 "github.com/Bifang-Bird/goOauth2"
)

// Client client model
type ClientPassword struct {
	ID        string
	Secret    string
	Domain    string
	Public    bool
	UserID    string
	Password  string
	Account   string
	GrantType oauth2.GrantType
}

// GetID client id
func (c *ClientPassword) GetID() string {
	return c.ID
}

// GetSecret client secret
func (c *ClientPassword) GetSecret() string {
	return c.Secret
}

// GetDomain client domain
func (c *ClientPassword) GetDomain() string {
	return c.Domain
}

// IsPublic public
func (c *ClientPassword) IsPublic() bool {
	return c.Public
}

// GetUserID user id
func (c *ClientPassword) GetUserID() string {
	return c.UserID
}

// GetUserID user id
func (c *ClientPassword) GetAccount() string {
	return c.Account
}

func (c *ClientPassword) GetGrantType() oauth2.GrantType {
	return c.GrantType
}

func (c *ClientPassword) VerifyPassword(password string) bool {
	booll := strings.EqualFold(c.Password, password)
	fmt.Printf("VerifyPassword c.password = %v , password = %v , bool= %v\n", c.Password, password, booll)
	return booll

}
