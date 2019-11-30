package responses

import "github.com/alidhkh/gista/models"

type TokenResult struct {
	Response
	Token models.Token
}
