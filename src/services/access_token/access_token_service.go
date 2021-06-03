package access_token

import (
	"github.com/juliandev/bookstore_oauth-api/src/domain/access_token"
	"github.com/juliandev/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
        GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpiresTime(access_token.AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
        UpdateExpiresTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(accessToken access_token.AccessToken) *errors.RestErr {
	if err := accessToken.Validate(); err != nil  {
		return err
	}
	return s.repository.Create(accessToken)
}

func (s *service) UpdateExpiresTime(accessToken access_token.AccessToken) *errors.RestErr {
	if err := accessToken.Validate(); err != nil  {
                return err
        }
        return s.repository.UpdateExpiresTime(accessToken)
}
