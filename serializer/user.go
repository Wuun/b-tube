package serializer

import "btube/model"

// User is the serialzer of user.
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// UserResponse is user response.
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser id the user response.
func BuildUser(user *model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse serialzer the user response.
func BuildUserResponse(user *model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
