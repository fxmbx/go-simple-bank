package token

import (
	"time"
)

//Maker is an interface for managing tokens
type Maker interface {
	//Cretae and send a token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//check if token is valid of not and return the payload data stored in the token
	VerifyToken(token string) (*Payload, error)
}
