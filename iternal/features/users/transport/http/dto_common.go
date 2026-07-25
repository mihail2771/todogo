package user_transport_http

import "github.com/mihail2771/todogo/iternal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id" example:"10"`
	Version     int     `json:"version" example:"3"`
	FullName    string  `json:"full_name" example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number" example:"+79994412345"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	userDTO := make([]UserDTOResponse, len(users))
	for i, user := range users {
		userDTO[i] = userDTOFromDomain(user)
	}

	return userDTO

}
