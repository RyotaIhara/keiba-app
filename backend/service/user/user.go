// Package user(Service)
package user

import (
	userModel "tmp-app-backend/model/user"
)

func GetUsers() []userModel.User {
	users := []userModel.User{
		{
			ID:   1,
			Code: "test001",
			Name: "テストユーザー001",
		},
		{
			ID:   2,
			Code: "test002",
			Name: "テストユーザー002",
		},
		{
			ID:   2,
			Code: "test002",
			Name: "テストユーザー002",
		},
	}

	return users
}
