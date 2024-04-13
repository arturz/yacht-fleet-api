package models

// token do operacji w trybie "POST once exactly"
type Token struct {
	ID   int  `json:"id"`
	used bool `json:"used"`
}

var Tokens = map[int]Token{}
var tokens_id_counter = 0

func NewToken() *Token {
	token := &Token{
		ID:   tokens_id_counter,
		used: false,
	}

	tokens_id_counter++
	Tokens[token.ID] = *token

	return token
}

func ClearTokens() {
	Tokens = map[int]Token{}
}

func IsTokenUsed(token_id int) bool {
	token, exists := Tokens[token_id]
	if !exists {
		return false
	}

	return token.used
}

func UseToken(token_id int) {
	token, exists := Tokens[token_id]
	if !exists {
		return
	}

	token.used = true
	Tokens[token_id] = token
}
