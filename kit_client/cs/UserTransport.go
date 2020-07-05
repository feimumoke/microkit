package cs

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func GetUserInfoRequest(_ context.Context, request *http.Request, r interface{}) error {
	userRequest := r.(UserRequest)
	request.URL.Path += "/user/" + strconv.Itoa(userRequest.Uid)
	return nil
}

func GetUserInfoResponse(_ context.Context, response *http.Response) (resp interface{}, err error) {
	if response.StatusCode > 400 {
		return nil, errors.New("no data")
	}
	var user_resp UserResponse
	err = json.NewDecoder(response.Body).Decode(&user_resp)
	if err != nil {
		return nil, err
	}
	return user_resp, nil
}
