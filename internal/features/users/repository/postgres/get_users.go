package users_postgres_repositoty

import (
	"context"
	"fmt"

	"github.com/mihail2771/todogo/iternal/core/domain"
)

func (r *UserRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	defer rows.Close()

	var usersModel []UserModel
	for rows.Next() {
		var userModel UserModel
		if err := rows.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.FullName,
			&userModel.PhoneNumber,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		usersModel = append(usersModel, userModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	usersDomain := UserDomainsFromModels(usersModel)

	return usersDomain, nil
}
