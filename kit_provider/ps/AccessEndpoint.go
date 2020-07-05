package ps

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	"time"
)

const secKey = "123abc"

type UserClaim struct {
	Uname string `json:"username"`
	jwt.StandardClaims
}

type IAccessService interface {
	GetToken(uname string, upass string) (string, error)
}
type AccessService struct {
}

type AccessRequest struct {
	Username string
	Userpass string
	Method   string
}

type AccessResponse struct {
	Status string
	Token  string
}

func (this *AccessService) GetToken(uname string, upass string) (string, error) {
	if upass == "123" {
		userInfo := &UserClaim{Uname: uname}
		userInfo.ExpiresAt = time.Now().Add(time.Second *220).Unix()
		token_obj := jwt.NewWithClaims(jwt.SigningMethodHS256, userInfo)
		token, err := token_obj.SignedString([]byte(secKey))
		return token, err
	}
	return "", fmt.Errorf("error username or password")
}

func AccessEndponit(ser IAccessService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AccessRequest)
		resp := AccessResponse{Status: "OK"}
		if req.Method == "POST" {
			token, err := ser.GetToken(req.Username, req.Userpass)
			if err != nil {
				resp.Status = "error:" + err.Error()
			} else {
				resp.Token=token
			}
			return resp,nil
		}
		return nil,errors.New("method not allow")
	}
}
