package responses

import "github.com/alidhkh/gista/models"

type Friendship struct {
	Response
	FriendshipStatus models.FriendshipStatus `json:"friendship_status"`
}
