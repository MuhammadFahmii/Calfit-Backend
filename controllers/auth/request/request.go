package request

import (
	"CalFit/business/superadmins"
	"CalFit/business/users"
)

type Auth struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SuperadminAuth struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (s SuperadminAuth) ToDomain() superadmins.Domain {
	return superadmins.Domain{
		Username: s.Username,
		Password: s.Password,
	}
}

func (a Auth) ToDomain() users.Domain {
	return users.Domain{
		Email:    a.Email,
		Password: a.Password,
	}
}
