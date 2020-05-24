package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func GetUserInfoRequest(_ context.Context, request *http.Request,r interface{})  error {
	userRequest := r.(UserRequest)
	request.URL.Path += "/user/" + strconv.Itoa(userRequest.Uid)
	return nil
}

func GetUserInfoResponse(_ context.Context, res *http.Response) (response interface{},err error) {
	if res.StatusCode >= 400{
		return nil,errors.New("no data")
	}
	var resp UserResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil{
		return nil,err
	}
	return resp,nil
}
