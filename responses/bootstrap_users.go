package responses

import "github.com/alidhkh/gista/models"

type BootstrapUsers struct {
	Response
	Surfaces []models.Surface `json:"surfaces"`
	Users    []models.User    `json:"users"`
}
