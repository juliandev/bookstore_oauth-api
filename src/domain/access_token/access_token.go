package access_token

import (
	"github.com/juliandev/bookstore_utils-go/rest_errors"
	"github.com/juliandev/bookstore_oauth-api/src/utils/crypto_utils"
	"fmt"
	"time"
	"strings"
)

const (
	expirationTime            = 24
	grantTypePassword         = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessToken) Validate() rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func (at *AccessTokenRequest) Validate() rest_errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return rest_errors.NewBadRequestError("invalid grand_type parameter")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken {
		UserId: userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Now().UTC().After(time.Unix(at.Expires, 0))
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
