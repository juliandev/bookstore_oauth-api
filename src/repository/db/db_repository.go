package db

import (
	"github.com/juliandev/bookstore_oauth-api/src/domain/access_token"
	"github.com/juliandev/bookstore_oauth-api/src/utils/errors"
	"github.com/juliandev/bookstore_oauth-api/src/clients/redis"
	"encoding/json"
	"context"
)

var (
	ctx = context.Background()
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string)(*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpiresTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {}

func (r *dbRepository) GetById(id string)(*access_token.AccessToken, *errors.RestErr) {
	session, err := redis.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	val, err := session.Get(ctx, id).Result()
	if err != nil {
		return nil, errors.NewNotFoundError("no access token found with given id")
	}
	var result access_token.AccessToken
	if err = json.Unmarshal([]byte(val), &result); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(accessToken access_token.AccessToken) *errors.RestErr {
	session, err := redis.GetSession()
        if err != nil {
                return errors.NewInternalServerError(err.Error())
        }
	if value, _ := r.GetById(accessToken.AccessToken); value != nil {
                return errors.NewConflictError("access token already exists")
        }
	value, err := json.Marshal(accessToken)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	err = session.Set(ctx, accessToken.AccessToken, value, 0).Err()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpiresTime(accessToken access_token.AccessToken) *errors.RestErr {
        session, err := redis.GetSession()
        if err != nil {
                return errors.NewInternalServerError(err.Error())
        }
	if _, err := r.GetById(accessToken.AccessToken); err != nil {
		return errors.NewNotFoundError("no access token found with given id")
	}
	value, err := json.Marshal(accessToken)
        if err != nil {
                return errors.NewInternalServerError(err.Error())
        }
        err = session.Set(ctx, accessToken.AccessToken, value, 0).Err()
        if err != nil {
                return errors.NewInternalServerError(err.Error())
        }
        return nil
}

