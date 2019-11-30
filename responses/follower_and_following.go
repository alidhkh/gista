package responses

import "github.com/alidhkh/gista/models"

type FollowerAndFollowing struct {
	Response
	Users                         []models.User         `json:"users"`
	SuggestedUsers                models.SuggestedUsers `json:"suggested_users"`
	TruncateFollowRequestsAtIndex int                   `json:"truncate_follow_requests_at_index"`
	NextMaxId                     string                `json:"next_max_id"`
	PageSize                      int64                 `json:"page_size"`
	BigList                       bool                  `json:"big_list"`
}
