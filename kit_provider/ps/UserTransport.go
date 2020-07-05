package ps

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"zhuhui.com/microkit/kit_provider/util"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	//if r.URL.Query().Get("uid") != "" {
	//	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	//	return UserRequest{
	//		Uid: uid,
	//	}, nil
	//}
	vars := mux.Vars(r)
	if uid, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{
			Uid:    uid,
			Method: r.Method,
			Token:  r.URL.Query().Get("token"),
		}, nil
	}
	return nil, errors.New("Params error")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-Type", contentType)
	if myerr, ok := err.(*util.MyError); ok {
		w.WriteHeader(myerr.Code)
		w.Write(body)
	} else {
		w.WriteHeader(500)
		w.Write(body)
	}
}
