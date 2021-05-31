package access_token

import (
	"github.com/juliandev/bookstore_oauth-api/src/domain/access_token"
	"github.com/juliandev/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
        GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
