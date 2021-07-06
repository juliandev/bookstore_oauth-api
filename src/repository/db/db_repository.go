package db

import (
	"github.com/juliandev/bookstore_oauth-api/src/domain/access_token"
	"github.com/juliandev/bookstore_utils-go/rest_errors"
	"github.com/juliandev/bookstore_oauth-api/src/clients/redis"
	"encoding/json"
	"context"
	"errors"
)

var (
	ctx = context.Background()
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string)(*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessToken) rest_errors.RestErr
	UpdateExpiresTime(access_token.AccessToken) rest_errors.RestErr
}

type dbRepository struct {}

func (r *dbRepository) GetById(id string)(*access_token.AccessToken, rest_errors.RestErr) {
	session := redis.GetSession()
	val, err := session.Get(ctx, id).Result()
	if err != nil {
		return nil, rest_errors.NewNotFoundError("no access token found with given id")
	}
	var result access_token.AccessToken
	if err = json.Unmarshal([]byte(val), &result); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to get current id", errors.New("database error"))
	}
	return &result, nil
}

func (r *dbRepository) Create(accessToken access_token.AccessToken) rest_errors.RestErr {
	session := redis.GetSession()
	if value, _ := r.GetById(accessToken.AccessToken); value != nil {
                return rest_errors.NewConflictError("access token already exists")
        }
	value, err := json.Marshal(accessToken)
	if err != nil {
		return rest_errors.NewInternalServerError("error unmarshaling json", errors.New("server error"))
	}
	err = session.Set(ctx, accessToken.AccessToken, value, 0).Err()
	if err != nil {
		return rest_errors.NewInternalServerError("error saving in database", errors.New("database error"))
	}
	return nil
}

func (r *dbRepository) UpdateExpiresTime(accessToken access_token.AccessToken) rest_errors.RestErr {
        session := redis.GetSession()
	if _, err := r.GetById(accessToken.AccessToken); err != nil {
		return rest_errors.NewNotFoundError("no access token found with given id")
	}
	value, err := json.Marshal(accessToken)
        if err != nil {
                return rest_errors.NewInternalServerError("error unmarshaling json", errors.New("server error"))
        }
        err = session.Set(ctx, accessToken.AccessToken, value, 0).Err()
        if err != nil {
                return rest_errors.NewInternalServerError("error when trying to update current resource", errors.New("database error"))
        }
        return nil
}

