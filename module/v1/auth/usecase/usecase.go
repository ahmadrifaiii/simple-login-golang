package usecase

import (
	"errors"
	"go-lemon/config"
	"go-lemon/model"
	"go-lemon/module/v1/user/repo"
	"go-lemon/utl/jwt"
	"go-lemon/utl/password"
)

func Login(cnf config.Configuration, p *model.LoginPayload) (token model.AuthAccess, err error) {

	// validation of null data
	if p.Email == "" || p.Password == "" {
		err = errors.New("email password can not null")
		return
	}

	// find user
	getUser, err := repo.GetUserDetailByParam(cnf.MysqlDB, "user_email", p.Email)
	if err != nil {
		return
	}

	if err = password.Decrypt([]byte(getUser.Password), p.Password); err != nil {
		return
	}

	return jwt.Generate(getUser.Email)
}
