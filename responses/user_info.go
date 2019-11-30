package responses

import "github.com/alidhkh/gista/models"

type UserInfo struct {
	Response
	Megaphone *interface{} `json:"megaphone,omitempty"`
	User      *models.User `json:"user,omitempty"`
}
