package user_transport_http

import "github.com/mihail2771/todogo/iternal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
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
