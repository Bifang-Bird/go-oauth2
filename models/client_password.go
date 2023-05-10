package models

import "strings"

// Client client model
type ClientPassword struct {
	ID       string
	Secret   string
	Domain   string
	Public   bool
	UserID   string
	Password string
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

func (c *ClientPassword) VerifyPassword(password string) bool {
	if len(password) <= 0 {
		return false
	}
	return strings.EqualFold(c.Password, password)

}
