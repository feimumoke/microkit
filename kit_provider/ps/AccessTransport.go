package ps

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func DecodeAccessRequest(c context.Context, r *http.Request) (interface{}, error) {
	body, _ := ioutil.ReadAll(r.Body)
	result := gjson.Parse(string(body))
	if result.IsObject() {
		username := result.Get("username")
		userpass := result.Get("userpass")
		return AccessRequest{Username: username.String(), Userpass: userpass.String(), Method: r.Method}, nil
	}
	return nil, errors.New("access param error")
}
func EncodeAccessResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Context-type", "application-json")
	return json.NewEncoder(w).Encode(response)
}
