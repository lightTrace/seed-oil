package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	if uid, ok := params["uid"]; ok {
		userId, _ := strconv.Atoi(uid)
		return UserRequest{Uid: userId, Method: r.Method}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
