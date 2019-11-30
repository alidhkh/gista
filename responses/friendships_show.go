package responses

import "github.com/alidhkh/gista/models"

type FriendshipsShow struct {
	Response
	models.FriendshipStatus
}
