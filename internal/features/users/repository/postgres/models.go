package users_postgres_repositoty

import "github.com/mihail2771/todogo/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func UserDomainsFromModels(userModels []UserModel) []domain.User {
	userDomains := make([]domain.User, len(userModels))
	for i, userModel := range userModels {
		userDomains[i] = domain.NewUser(
			userModel.ID,
			userModel.Version,
			userModel.FullName,
			userModel.PhoneNumber,
		)
	}
	return userDomains

}
