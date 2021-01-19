package oauth

import (
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/sshindanai/repo/bookstore-users-api/domain/usersdomain"
	"github.com/sshindanai/repo/bookstore-users-api/rest/models"
	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:1323",
		Timeout: 100 * time.Millisecond,
	}
)

const (
	expirationTime            = 24
	grantTypePassword         = "password"
	grantTypeClientCredential = "client_credentials"
)

type RestOauth interface {
	Authenticate(*usersdomain.User) (*models.AccessToken, *errors.RestErr)
}

type restOauth struct{}

func NewRestOauth() RestOauth {
	return &restOauth{}
}

func (r *restOauth) Authenticate(user *usersdomain.User) (*models.AccessToken, *errors.RestErr) {
	request := models.AccessTokenRequest{
		Username:  user.Email,
		Password:  user.Password,
		GrantType: grantTypePassword,
	}

	response := oauthRestClient.Post("/oauth/access_token", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewUnauthorizedError("invalid restclient response when trying to login user (users side)")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		if err := json.Unmarshal(response.Bytes(), &restErr); err != nil {
			// Happen when tag "status_code" is string
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var at models.AccessToken
	if err := json.Unmarshal(response.Bytes(), &at); err != nil {
		return nil, errors.NewUnauthorizedError("error when trying to unmarshal users login response")
	}
	return &at, nil
}
