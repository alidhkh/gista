package responses

import "github.com/alidhkh/gista/models"

type CreateBusinessInfo struct {
	Response
	Users []models.User `json:"users,omitempty"`
}
