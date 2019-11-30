package responses

import "github.com/alidhkh/gista/models"

type Comment struct {
	Response
	Comment *models.Comment `json:"comment"`
}
