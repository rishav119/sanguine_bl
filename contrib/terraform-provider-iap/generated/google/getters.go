package google

import "golang.org/x/oauth2"

// GetTokenSource exports the token source for use by the provider
func (c *Config) GetTokenSource() oauth2.TokenSource {
	return c.tokenSource
}
