package rest

import (
	"github.com/juliandev/bookstore_oauth-api/src/utils/errors"
	"github.com/juliandev/bookstore_oauth-api/src/domain/users"
	"github.com/go-resty/resty/v2"
)

var (
	usersRestClient = resty.New().R()
)

const (
	BaseURL = "http://localhost:8080"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {}

func NewRepository() RestUsersRepository {
	return &usersRepository
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	response, err := usersRestClient.SetHeader("Content-Type", "application/json").SetBody(request).Post(BaseURL + "/users/login")
	if err != nil {
		return nil, errors.RestErr(err.Error())
	}
	if response.Status() > 299 {
		var restErr errors.RestErr
		err =:= json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil

}
