package models

// Client client model
type ClientPermission struct {
	ClientID  string
	UserID    string
	ApiModule string
	ApiUrl    string
}

// GetID client id
func (c *ClientPermission) GetClientId() string {
	return c.ClientID
}

// GetSecret client secret
func (c *ClientPermission) GetUserID() string {
	return c.UserID
}

// GetDomain client domain
func (c *ClientPermission) GetApiModule() string {
	return c.ApiModule
}

// IsPublic public
func (c *ClientPermission) GetApiUrl() string {
	return c.ApiUrl
}
