package access_token

import (
	"github.com/juliandev/bookstore_oauth-api/src/domain/access_token"
	"github.com/juliandev/bookstore_oauth-api/src/repository/rest"
	"github.com/juliandev/bookstore_utils-go/rest_errors"
	"strings"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessToken) rest_errors.RestErr
	UpdateExpiresTime(access_token.AccessToken) rest_errors.RestErr
}

type Service interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr)
        UpdateExpiresTime(access_token.AccessToken) rest_errors.RestErr
}

type service struct {
	restUsersRepository rest.RestUsersRepository
	repository Repository
}

func NewService(repo Repository, usersRepo rest.RestUsersRepository) Service {
	return &service{
		restUsersRepository: usersRepo,
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, rest_errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.restUsersRepository.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	if err := s.repository.Create(at); err != nil {
		return nil, err
	}
	return &at, nil

}

func (s *service) UpdateExpiresTime(accessToken access_token.AccessToken) rest_errors.RestErr {
	if err := accessToken.Validate(); err != nil  {
                return err
        }
        return s.repository.UpdateExpiresTime(accessToken)
}
