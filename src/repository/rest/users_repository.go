package rest

import (
	"github.com/juliandev/bookstore_utils-go/rest_errors"
	"github.com/juliandev/bookstore_oauth-api/src/domain/users"
	"github.com/go-resty/resty/v2"
	"encoding/json"
	"errors"
	"os"
)

var (
	baseURL		string
	usersApiHost    = os.Getenv("USERS_API_HOST")
	usersRestClient = resty.New().R()
)

const (
	usersApiLocal = "http://localhost:8080"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct {}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest {
		Email: email,
		Password: password,
	}

	if usersApiHost == "" {
		baseURL = usersApiLocal
	} else {
		baseURL = usersApiHost
	}

	response, err := usersRestClient.SetHeader("Content-Type", "application/json").SetBody(request).Post(baseURL + "/users/login")
	if err != nil {
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying to login user", errors.New("restclient error"))
	}
	if response.StatusCode() > 299 {
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
	}
	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users response", errors.New("json parsing error"))
	}
	return &user, nil
}
