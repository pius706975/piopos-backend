package user

import "github.com/pius-microservices/piopos-database/database/models/user"

var RoleSeed = user.Roles{
	{
		ID: "8caed36f-9811-4e66-abd9-3b359d41a0c4",
		Name: "piopos-admin",
	},

	{
		ID: "e10e672b-4e59-495d-b0c6-1083b038832f",
		Name: "business-owner",
	},
}