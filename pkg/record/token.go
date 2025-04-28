package record

import "time"

// Permissions represents token permissions
type Permissions struct {
	FullAccess bool     // full access to manage buckets and tokens
	Read       []string // list of buckets with read access
	Write      []string // list of buckets with write access
}

// Token represents a token for authentication
type Token struct {
	Name          string    // name of token
	CreatedAt     time.Time // creation time of token
	IsProvisioned bool      // token is provisioned and can't be deleted or changed
}

// FullTokenInfo represents full information about a token with permissions
type FullTokenInfo struct {
	*Token
	Permissions Permissions // permissions of token
}

// TokenList represents a list of tokens
type TokenList struct {
	Tokens []Token // list of tokens
}

// TokenCreateResponse represents the response from creating a token
type TokenCreateResponse struct {
	Value string // token for authentication
}

// GetTokenList returns a list of tokens
func (t *Token) GetTokenList() ([]Token, error) {
	// make http request to get the list of tokens

	return []Token{}, nil
}

// GetToken returns a token
func (t *Token) GetToken() (Token, error) {
	// make http request to get token

	return Token{}, nil
}

// GenerateToken returns a token
func (t *Token) GenerateToken() (Token, error) {
	// make http request to generate a token

	return Token{}, nil
}

// RemoveToken removes a token
func (t *Token) RemoveToken(name string) error {
	// make http request to remove a token

	return nil
}
