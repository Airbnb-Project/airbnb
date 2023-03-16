package handler

import "airbnb/feature/user"

type UserResponse struct {
	ID             uint   `json:"id" form:"id"`
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	Phone          string `json:"phone" form:"phone"`
	Address        string `json:"address" form:"address"`
	Role           string `json:"role" form:"role"`
	ProfilePicture string `json:"photo" form:"photo"`
}

func LoginResponse(data user.Core) UserResponse {
	return UserResponse{
		ID:      data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Phone:   data.Phone,
		Address: data.Address,
		Role:    data.Role,
	}
}

func ProfileResponse(data user.Core) UserResponse {
	return UserResponse{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		Phone:          data.Phone,
		Address:        data.Address,
		Role:           data.Role,
		ProfilePicture: data.ProfilePicture,
	}
}
