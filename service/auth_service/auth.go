package auth_service

import models "github.com/ypMarkJo/Go-gin-samplemodels"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}
