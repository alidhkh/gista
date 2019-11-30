package responses

import "github.com/alidhkh/gista/models"

type RecentSearches struct {
	Response
	Recent []models.Suggested `json:"recent"`
}
