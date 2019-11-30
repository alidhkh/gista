package responses

import "github.com/alidhkh/gista/models"

type SwitchPersonalProfile struct {
	Response
	Users []models.User `json:"users,omitempty"`
}
