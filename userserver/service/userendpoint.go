package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"gokit/userserver/util"
	"strconv"
)

type UserRequest struct {
	Uid int `json:"uid"`
	Method string
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEndpoint(userService IUserService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := "nothing"
		if r.Method == "GET"{
			result = userService.GetName(r.Uid) + strconv.Itoa(util.ServicePort)
		}
		if r.Method == "DELETE"{
			err := userService.DelName(r.Uid)
			if err != nil{
				result = err.Error()
			}else{
				result = fmt.Sprintf("userId为%d的用户删除成功",r.Uid)
			}
		}
		return UserResponse{Result:result},nil
	}
}
